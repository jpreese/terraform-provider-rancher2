package rancher2

import (
	"github.com/hashicorp/terraform/helper/schema"
	managementClient "github.com/rancher/types/client/management/v3"
)

// Flatteners

func flattenGlobalRoleBinding(d *schema.ResourceData, in *managementClient.GlobalRoleBinding) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("global_role_id", in.GlobalRoleID)
	d.Set("user_id", in.UserID)
	d.Set("name", in.Name)

	err := d.Set("annotations", toMapInterface(in.Annotations))
	if err != nil {
		return err
	}

	err = d.Set("labels", toMapInterface(in.Labels))
	if err != nil {
		return err
	}

	return nil

}

// Expanders

func expandGlobalRoleBinding(in *schema.ResourceData) *managementClient.GlobalRoleBinding {
	obj := &managementClient.GlobalRoleBinding{}
	if in == nil {
		return nil
	}

	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.GlobalRoleID = in.Get("global_role_id").(string)
	obj.UserID = in.Get("user_id").(string)
	obj.Name = in.Get("name").(string)

	if v, ok := in.Get("annotations").(map[string]interface{}); ok && len(v) > 0 {
		obj.Annotations = toMapString(v)
	}

	if v, ok := in.Get("labels").(map[string]interface{}); ok && len(v) > 0 {
		obj.Labels = toMapString(v)
	}

	return obj
}
