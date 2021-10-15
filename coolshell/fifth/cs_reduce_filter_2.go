package fifth

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/8 11:12
 * @description: TODO
 */

type mysql_server struct {
	host string
	port uint32
}

type phyresource struct {
	cpuCount uint8
	memSize  uint8
	diskSize uint8
}

type mysql_instance struct {
	clusterId      uint32
	clusterServers []mysql_server
	databases      []string
	status         string
	resource       phyresource
}

var instances = []mysql_instance{
	{1, []mysql_server{{"192.168.1.100", 3306}, {"192.168.1.101", 3306}}, []string{"aaa", "bbb"}, "up",
		phyresource{2, 16, 100}},
	{2, []mysql_server{{"192.168.1.102", 3306}, {"192.168.1.103", 3306}}, []string{"ccc", "ddd", "eee"}, "up",
		phyresource{4, 32, 100}},
	{3, []mysql_server{{"192.168.1.104", 3306}, {"192.168.1.105", 3306}}, []string{"fffffff", "ee", "ggggg"}, "up",
		phyresource{6, 24, 300}},
	{4, []mysql_server{{"192.168.1.106", 3306}, {"192.168.1.107", 3306}}, []string{"hhh", "iii", "jjjj", "kkkk",
		"llll"},
		"down",
		phyresource{2, 16, 600}},
}

func instanceCountIf(mi []mysql_instance, f func(m *mysql_instance) bool) int {
	count := 0
	for i, _ := range mi {
		if f(&mi[i]) {
			count += len(mi[i].databases)
		}
	}
	return count
}

func instanceFilterIn(mi []mysql_instance, f func(m *mysql_instance) bool) []string {
	resmi := []string{}
	for i, _ := range mi {
		if f(&mi[i]) {
			resmi = append(resmi, mi[i].databases...)
		}
	}
	return resmi
}

func instanceSumOf(mi []mysql_instance, f func(m *mysql_instance) int) int {
	count := 0
	for i, _ := range mi {
		count += f(&mi[i])
	}
	return count
}
