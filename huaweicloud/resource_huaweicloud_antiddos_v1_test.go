package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/huaweicloud/golangsdk/openstack/antiddos/v1/antiddos"
)

func TestAccAntiDdosV1_basic(t *testing.T) {
	var antiddos antiddos.GetResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAntiDdosV1Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAntiDdosV1_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAntiDdosV1Exists("huaweicloud_antiddos.antiddos_1", &antiddos),
					resource.TestCheckResourceAttr(
						"huaweicloud_antiddos.antiddos_1", "enable_l7", "true"),
					resource.TestCheckResourceAttr(
						"huaweicloud_antiddos.antiddos_1", "traffic_pos_id", "1"),
					resource.TestCheckResourceAttr(
						"huaweicloud_antiddos.antiddos_1", "http_request_pos_id", "3"),
					resource.TestCheckResourceAttr(
						"huaweicloud_antiddos.antiddos_1", "cleaning_access_pos_id", "1"),
					resource.TestCheckResourceAttr(
						"huaweicloud_antiddos.antiddos_1", "app_type_id", "0"),
				),
			},
			// {
			// 	Config: testAccAntiDdosV1_update,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr(
			// 			"huaweicloud_antiddos.antiddos_1", "traffic_pos_id", "2"),
			// 		resource.TestCheckResourceAttr(
			// 			"huaweicloud_antiddos.antiddos_1", "http_request_pos_id", "1"),
			// 		resource.TestCheckResourceAttr(
			// 			"huaweicloud_antiddos.antiddos_1", "cleaning_access_pos_id", "2"),
			// 		resource.TestCheckResourceAttr(
			// 			"huaweicloud_antiddos.antiddos_1", "app_type_id", "1"),
			// 	),
			// },
		},
	})
}

func TestAccAntiDdosV1_timeout(t *testing.T) {
	var antiddos antiddos.GetResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAntiDdosV1Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAntiDdosV1_timeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAntiDdosV1Exists("huaweicloud_antiddos.antiddos_1", &antiddos),
				),
			},
		},
	})
}

func testAccCheckAntiDdosV1Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	antiddosClient, err := config.antiddosV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating antiddos client: %s", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_antiddos" {
			continue
		}
		n, err := antiddos.Get(antiddosClient, "ee73fd6b-ee18-4d29-a077-bbe31502eaa4").Extract()
		if err == nil {
			fmt.Printf(" ***** n = %v", n)
			return fmt.Errorf("antiddos still exists")
		}
	}

	return nil
}

func testAccCheckAntiDdosV1Exists(n string, ddos *antiddos.GetResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		antiddosClient, err := config.antiddosV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating antiddos client: %s", err)
		}

		found, err := antiddos.Get(antiddosClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		*ddos = *found

		return nil
	}
}

const testAccAntiDdosV1_basic = `
resource "huaweicloud_vpc_eip_v1" "eip_1" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name = "test_antiddos3"
    size = 8
    share_type = "PER"
    charge_mode = "traffic"
  }
}

resource "huaweicloud_antiddos" "antiddos_1" {
  floating_ip_id = "${huaweicloud_vpc_eip_v1.eip_1.id}"
  enable_l7 = true
  traffic_pos_id = 1
  http_request_pos_id = 3
  cleaning_access_pos_id = 1
  app_type_id = 0
}
`
const testAccAntiDdosV1_update = `
resource "huaweicloud_vpc_eip_v1" "eip_1" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name = "test_antiddos3"
    size = 8
    share_type = "PER"
    charge_mode = "traffic"
  }
}

resource "huaweicloud_antiddos" "antiddos_1" {
  floating_ip_id = "${huaweicloud_vpc_eip_v1.eip_1.id}"
  enable_l7 = true
  traffic_pos_id = 2
  http_request_pos_id = 1
  cleaning_access_pos_id = 2
  app_type_id = 1
}
`

const testAccAntiDdosV1_timeout = `
resource "huaweicloud_vpc_eip_v1" "eip_1" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name = "test_antiddos"
    size = 8
    share_type = "PER"
    charge_mode = "traffic"
  }
}

resource "huaweicloud_antiddos" "antiddos_1" {
  floating_ip_id = "${huaweicloud_vpc_eip_v1.eip_1.id}"
  enable_l7 = true
  traffic_pos_id = 1
  http_request_pos_id = 2
  cleaning_access_pos_id = 1
  app_type_id = 0

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`
