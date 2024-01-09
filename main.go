package main

import enumValueMap "github.com/valllabh/ocsf-schema-golang/ocsf/v1_0_0/enum-value-map"

// Nothing to see here, move along

func main() {

}

func testEnumValue() {
	ev, _ := enumValueMap.GetEnumValue("ACCOUNT_CHANGE_CATEGORY_UID_IDENTITY_ACCESS_MANAGEMENT")
	println(ev.Name, ev.Value)
}
