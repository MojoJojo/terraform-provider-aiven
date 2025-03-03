// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/
package aiven

import (
	"context"
	"fmt"
	"strings"

	"github.com/aiven/aiven-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var aivenMirrorMakerReplicationFlowSchema = map[string]*schema.Schema{
	"project": {
		Type:         schema.TypeString,
		Required:     true,
		Description:  "Project to link the kafka topic to",
		ForceNew:     true,
		ValidateFunc: validation.StringLenBetween(1, 63),
	},
	"service_name": {
		Type:         schema.TypeString,
		Required:     true,
		Description:  "Service to link the kafka topic to",
		ForceNew:     true,
		ValidateFunc: validation.StringLenBetween(1, 63),
	},
	"enable": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Enable of disable replication flows for a service",
	},
	"source_cluster": {
		Type:         schema.TypeString,
		Required:     true,
		Description:  "Source cluster alias",
		ValidateFunc: validation.StringLenBetween(1, 128),
	},
	"target_cluster": {
		Type:         schema.TypeString,
		Required:     true,
		Description:  "Target cluster alias",
		ValidateFunc: validation.StringLenBetween(1, 128),
	},
	"topics": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of topics and/or regular expressions to replicate",
		Elem: &schema.Schema{
			Type:     schema.TypeString,
			MaxItems: 256,
		},
	},
	"topics_blacklist": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of topics and/or regular expressions to not replicate.",
		Elem: &schema.Schema{
			Type:     schema.TypeString,
			MaxItems: 256,
		},
	},
	"replication_policy_class": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Replication policy class",
		Default:     "org.apache.kafka.connect.mirror.DefaultReplicationPolicy",
		ValidateFunc: validation.StringInSlice([]string{
			"org.apache.kafka.connect.mirror.DefaultReplicationPolicy",
			"org.apache.kafka.connect.mirror.IdentityReplicationPolicy"}, false),
	},
	"sync_group_offsets_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Sync consumer group offsets",
	},
	"sync_group_offsets_interval_seconds": {
		Type:         schema.TypeInt,
		Optional:     true,
		Description:  "Frequency of consumer group offset sync",
		ValidateFunc: validation.IntAtLeast(1),
		Default:      1,
	},
	"emit_heartbeats_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Emit heartbeats enabled",
	},
}

func resourceMirrorMakerReplicationFlow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMirrorMakerReplicationFlowCreate,
		ReadContext:   resourceMirrorMakerReplicationFlowRead,
		UpdateContext: resourceMirrorMakerReplicationFlowUpdate,
		DeleteContext: resourceMirrorMakerReplicationFlowDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceMirrorMakerReplicationFlowState,
		},

		Schema: aivenMirrorMakerReplicationFlowSchema,
	}
}

func resourceMirrorMakerReplicationFlowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	project := d.Get("project").(string)
	serviceName := d.Get("service_name").(string)
	enable := d.Get("enable").(bool)
	sourceCluster := d.Get("source_cluster").(string)
	targetCluster := d.Get("target_cluster").(string)

	err := client.KafkaMirrorMakerReplicationFlow.Create(project, serviceName, aiven.MirrorMakerReplicationFlowRequest{
		ReplicationFlow: aiven.ReplicationFlow{
			Enabled:                         enable,
			SourceCluster:                   sourceCluster,
			TargetCluster:                   d.Get("target_cluster").(string),
			Topics:                          flattenToString(d.Get("topics").([]interface{})),
			TopicsBlacklist:                 flattenToString(d.Get("topics_blacklist").([]interface{})),
			ReplicationPolicyClass:          d.Get("replication_policy_class").(string),
			SyncGroupOffsetsEnabled:         d.Get("sync_group_offsets_enabled").(bool),
			SyncGroupOffsetsIntervalSeconds: d.Get("sync_group_offsets_interval_seconds").(int),
			EmitHeartbeatsEnabled:           d.Get("emit_heartbeats_enabled").(bool),
		},
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildResourceID(project, serviceName, sourceCluster, targetCluster))

	return resourceMirrorMakerReplicationFlowRead(ctx, d, m)
}

func resourceMirrorMakerReplicationFlowRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	project, serviceName, sourceCluster, targetCluster := splitResourceID4(d.Id())
	replicationFlow, err := client.KafkaMirrorMakerReplicationFlow.Get(project, serviceName, sourceCluster, targetCluster)
	if err != nil {
		return diag.FromErr(resourceReadHandleNotFound(err, d))
	}

	if err := d.Set("project", project); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("service_name", serviceName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enable", replicationFlow.ReplicationFlow.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("source_cluster", sourceCluster); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_cluster", targetCluster); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("topics", replicationFlow.ReplicationFlow.Topics); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("topics_blacklist", replicationFlow.ReplicationFlow.TopicsBlacklist); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("replication_policy_class", replicationFlow.ReplicationFlow.ReplicationPolicyClass); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("sync_group_offsets_enabled", replicationFlow.ReplicationFlow.SyncGroupOffsetsEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("sync_group_offsets_interval_seconds", replicationFlow.ReplicationFlow.SyncGroupOffsetsIntervalSeconds); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("emit_heartbeats_enabled", replicationFlow.ReplicationFlow.EmitHeartbeatsEnabled); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceMirrorMakerReplicationFlowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	project, serviceName, sourceCluster, targetCluster := splitResourceID4(d.Id())
	_, err := client.KafkaMirrorMakerReplicationFlow.Update(
		project,
		serviceName,
		sourceCluster,
		targetCluster,
		aiven.MirrorMakerReplicationFlowRequest{
			ReplicationFlow: aiven.ReplicationFlow{
				Enabled:                         d.Get("enable").(bool),
				Topics:                          flattenToString(d.Get("topics").([]interface{})),
				TopicsBlacklist:                 flattenToString(d.Get("topics_blacklist").([]interface{})),
				ReplicationPolicyClass:          d.Get("replication_policy_class").(string),
				SyncGroupOffsetsEnabled:         d.Get("sync_group_offsets_enabled").(bool),
				SyncGroupOffsetsIntervalSeconds: d.Get("sync_group_offsets_interval_seconds").(int),
				EmitHeartbeatsEnabled:           d.Get("emit_heartbeats_enabled").(bool),
			},
		},
	)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceMirrorMakerReplicationFlowRead(ctx, d, m)
}

func resourceMirrorMakerReplicationFlowDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	project, serviceName, sourceCluster, targetCluster := splitResourceID4(d.Id())

	err := client.KafkaMirrorMakerReplicationFlow.Delete(project, serviceName, sourceCluster, targetCluster)
	if err != nil {
		diag.FromErr(err)
	}

	return nil
}

func resourceMirrorMakerReplicationFlowState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if len(strings.Split(d.Id(), "/")) != 4 {
		return nil, fmt.Errorf("invalid identifier %v, expected <project_name>/<service_name>/<source_cluster>/<target_cluster>", d.Id())
	}

	di := resourceMirrorMakerReplicationFlowRead(ctx, d, m)
	if di.HasError() {
		return nil, fmt.Errorf("cannot get mirror maker 2 replication flow: %v", di)
	}

	return []*schema.ResourceData{d}, nil
}
