package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
)

func ToJson(a any) string {
	b, _ := json.Marshal(a)
	return string(b)
}

type A struct {
	AA int    `json:"aaa"`
	BB string `json:"bbb"`
}

func test1() {
	err := xerrors.Errorf(": %w ", "ERROR")
	fmt.Printf("%v", err)

	str := `
	{
		"aaa":2,
		"ccc":3,
		"bbb":"4"
	}
	`
	var a A
	json.Unmarshal([]byte(str), &a)
	fmt.Printf("%+v", a)
}

func test2(a ...bool) {
	if len(a) == 1 && a[0] {
		fmt.Printf("TRUE")
	}
}

func test3() {
	str := fmt.Sprintf("%v : %v", 2, 3)
	fmt.Printf("%v, %T", str, str)
}
func test4() {
	st := "24242424244242" + fmt.Sprintf("temp_no_dup:cmp_lowest_price:%v:%v:%v:%v", 1, 2, 3, 4)
	fmt.Print(st)
}
func test5() {
	type PP struct {
		Test bool `json:"test_test"`
	}

	testStr := `{
		"test_test": true
	}`
	p := &PP{}
	pp := &p
	ppp := &pp
	json.Unmarshal([]byte(testStr), ppp)
	fmt.Printf("%+v", (*(*ppp)).Test)

}

func test6() error {
	return nil
}

func test7() {
	if xerrors.Errorf("%w", test6()) == nil {
		println(1)
	} else {
		println(2)
	}
}
func test8() {
	type MP struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
	}

	var mp = map[string]interface{}{
		"name": "123",
		"sex":  "55",
	}

	st, _ := json.Marshal(&MP{
		Name: "123",
		Sex:  "55",
	})

	m, _ := json.Marshal(mp)

	fmt.Print(st, "\n", m)
	print(string(st) == string(m))
}
func test9() {
	var mp = map[string]struct{}{}
	mp["keep"] = struct{}{}

	if _, ok := mp["keep"]; ok {
		print("ok1")
	}
}

func test10(arr []int64) {
	if len(arr) == 0 {
		return
	}
	arr = []int64{
		1, 2, 3,
	}
	fmt.Print(arr)
}
func test11() {
	arr := []int64{
		2, 2, 2,
	}
	test10(arr)
	fmt.Print(arr)
}
func test12() {
	type KK struct {
		Value int
		Sku   int
	}
	var a = map[int64]*KK{
		1: &KK{
			Value: 123,
			Sku:   122,
		},
	}

	if _, ok := a[2]; ok {
		fmt.Print(a[2])
	}

	fmt.Print(a[1])
}

func test15() {
	a := []int64{1, 2, 3}
	k := 1
	n := 2
	var x int64
	x, k = a[k], n

	print(x, k)
}

type PeopleService struct{}

// methods

func (p *PeopleService) method1() error {
	return nil
}

func (p *PeopleService) method2() {

}

func test18() {
	type B struct{}
	type A struct {
		a int
		b *B
	}

	var aa = &A{
		a: 2,
		b: &B{},
	}
	var aaaa = &A{
		a: 2,
		b: &B{},
	}

	if reflect.DeepEqual(aa, aaaa) {
		print("aaa")
	}
}

func test19() {
	type User struct {
		Name string
		Id   int
	}
	type S struct {
		A int
		B string
		C *User
	}

	v1 := &S{
		A: 1,
		B: "abc",
		C: &User{
			Name: "name",
			Id:   2222,
		},
	}

	fmt.Printf("v1:%v", *v1)
}
func ConvertStringOptionToInt64Slice(str string, option string) []int64 {
	var result []int64
	strs := strings.Split(str, option)
	for _, item := range strs {
		result = append(result, StringToInt64(item))
	}

	return result
}

func StringToInt64(keyword string) int64 {
	num, err := strconv.ParseInt(strings.TrimSpace(keyword), 10, 64)
	if err != nil {
		return 0
	} else {
		return num
	}
}

func test22() {
	var skuIds = "12,23121,1231313123,12313123123132,222"
	result := ConvertStringOptionToInt64Slice(skuIds, ",")

	fmt.Print(fmt.Sprint(result))
}

func test23() {
	// var a int = 3
	var x *int = nil

	print(*x)
}

func Del(a []int64, k int64) []int64 {
	var res = []int64{}
	for _, num := range a {
		if k != num {
			res = append(res, num)
		}
	}
	return res
}

func test24() {
	var r = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r = Del(r, 2)
	r = Del(r, 3)
	r = Del(r, 7)
	for _, i := range r {
		println(i)
	}
}

func test25() []string {
	var v map[string][]string = map[string][]string{
		"a": {
			"aa", "aaa",
		},
		"b": {
			"vbv", "aaavvv",
		},
	}

	return v["c"]
}

func test26() {
	var v = map[string]struct{}{}
	v["2"] = struct{}{}
	for k := range v {
		print(k)
	}
}

func test27() {
	type A struct {
		Id   int
		Name string
	}
	a := []A{
		{
			Id:   1,
			Name: "a",
		},
		{
			Id:   2,
			Name: "b",
		},
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Id > a[j].Id
	})

	fmt.Printf("%+v", a[0])
}

func test28() {
	type B struct {
		Key string
	}
	type A struct {
		Id   int
		Name string
		Bob  *B
	}
	var r = []*A{
		{
			Id:   1,
			Name: "1",
			Bob: &B{
				Key: "key1",
			},
		},
		{
			Id:   2,
			Name: "2",
			Bob: &B{
				Key: "key2",
			},
		},
		{
			Id:   3,
			Name: "3",
			Bob: &B{
				Key: "key3",
			},
		},
	}
	Fn := func(arrs []*A) []*A {
		var result []*A
		for _, arr := range arrs {
			if arr.Bob.Key == "key2" {
				result = append(result, arr)
			}
		}
		return result
	}

	r = Fn(r)
	for _, rr := range r {
		fmt.Println(rr)
	}
}

// 测试Marshal是否排序
func test29() {
	var a = map[string]interface{}{
		"123": 1,
		"133": 2,
		"22":  11,
		"100": 13,
	}

	b, _ := json.Marshal(a)
	print(string(b))
}

func test30() {
	type B struct {
		Name string
	}

	type A struct {
		Name string
		BBB  *B
	}

	a := &A{
		Name: "123",
		BBB: &B{
			Name: "12313213123",
		},
	}

	b := a

	b.BBB.Name = "123"

	fmt.Println(b == a)
	fmt.Println(&b == &a)
	fmt.Println(b.BBB.Name == a.BBB.Name)
}

func test31() {
	var a = []int64{1, 2, 3, 4, 5, 6}
	var b = a[1:]
	var c = a[2:]
	if b[1] == c[0] {
		fmt.Println(111)
		fmt.Println(a)
	}

	b[1] = 200
	c[0] = 300

	if b[1] == c[0] {
		fmt.Println(2222)
		fmt.Println(a)
	}
}

func test32() {
	go func() {
		var n = 1
		for {
			fmt.Printf("爱小妍的第%v天\n", n)

			time.Sleep(time.Second * 1)
			n++
		}
	}()

	var a = 0
	fmt.Scanf("%d", &a)

	fmt.Printf("a = %v\n \n --- DONE --- \n", a)
}

func test33() {
	c := make(chan int)

	go func() {
		for i := 2; i < 30; i++ {
			c <- i
			fmt.Println("发送到", i)
		}
	}()

	// for data := range c {
	// 	fmt.Println("接收到", data)
	// 	if data == 0 {

	// 		break
	// 	}
	// }
	var data = <-c
	fmt.Println("接收到", data)
	time.Sleep(time.Second)
	// data = <-c
	// fmt.Println("接收到", data)
	// data = <-c
	// fmt.Println("接收到", data)

}

func test34() {
	var a string = "我是你asdbv"
	for _, item := range a {
		fmt.Printf("%v\n", string(item))
	}
}

func test35() {
	t := time.Now().Unix()

	tt := time.Now().Add(-30 * 24 * time.Hour).Unix()

	fmt.Println(t, tt)
}

func test36() {
	type A struct {
		Name int64
		Str  int64
	}

	a := &A{}

	fmt.Println(a.Name, a.Str)
}
func test37() {
	var a中国变量 = 6666
	var b中国标识符 = "无敌的"

	fmt.Println(a中国变量, b中国标识符)
}

func test38() {
	type A struct {
		Name string
		Id   string
	}

	var a *A = &A{
		Name: "123",
		Id:   "123",
	}

	var b = a

	a = &A{
		Name: "1222",
		Id:   "1231321",
	}

	fmt.Println(b, a)
}

func test_between_string_and_fmtSprint() {
	var a uint = 213123132
	fmt.Printf("a: %v\n", a)

	t := fmt.Sprint(time.Now().Unix())

	fmt.Println(string(a))

	fmt.Println(fmt.Sprint(a))

	fmt.Println(t + string(a))

	fmt.Println(t + fmt.Sprint(a))

}

type AA struct {
}

func (a *AA) Get() int64 {
	return 1
}
func test_nil_struct_invoke_its_method() {
	var a *AA
	if a != nil {
		fmt.Print("123213123++", a.Get())
	}
	fmt.Print(a.Get(), a)
}

type Entity struct {
	Name string
	Id   int64
}

type DbContext struct {
	InitFactory     map[int64]*Entity
	UltimateFactory map[int64]*Entity
}

func (p *DbContext) SaveChanges() {
	// TODO: Save to DataBase
}

func test_context_reference_address() {

}

func test_sync_once() {
	// 创建一个 Once 对象
	once := sync.Once{}

	// 多次调用函数，只会执行一次
	for i := 0; i < 5; i++ {
		once.Do(func() {
			// 执行被包装的函数
			fmt.Println("执行函数:", i)
		})
	}

	// 释放 Once 对象
	once.Do(func() {
		fmt.Println("释放锁:", time.Now())
	})
}

func test_map_nil_access() {
	var a map[string]string
	if a == nil {
		fmt.Println(123)
	}
	fmt.Print(a["should"] == "true")
}

func test_time_format() {
	var t string = "2023-03-06 00:00:00"
	f, _ := time.Parse(t, "YYYY-MM-DD HH:mm:ss USTC")
	fmt.Println(f)
}

func SendAndReturn(a []int64) []int64 {
	return a
}

func test_int_array_return() {
	var arr = []int64{1, 2, 3, 4}
	arr1 := SendAndReturn(arr)

	arr1 = append(arr1, 5)
	fmt.Println(arr1)
	fmt.Println(arr)
}

func test_slice_index() {
	a := []int64{1, 2, 3, 4, 6, 7, 8, 9, 0}

	fmt.Print(a[2:9])
}

func TestMarshalSequencely() {
	var m = map[string]interface{}{
		"1":         3,
		"123123123": 1,
		"123123113": 2,
		"123":       3,
		"11":        4,
	}

	r, _ := json.Marshal(m)

	fmt.Print(string(r))
}

func TestTwoDimsArraySlice() {
	k := [][]int64{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(k[2:])
}

func GenerateClosure1() func(name string) string {
	var name string
	return func(n string) string {
		name = n
		return name
	}
}

type clazz struct {
	Name string `json:"name"`
}

func GenerateClosure2() func(name clazz) clazz {
	var c clazz
	return func(n clazz) clazz {
		c = n
		return c
	}
}

func TestClosureUse() {
	f1 := GenerateClosure1()
	f2 := GenerateClosure2()

	fmt.Println(f1("123"))
	fmt.Println(f2(clazz{Name: "123"}))
}

// callback
func Process() func() (int64, error) {
	var a, b, c int64
	// 处理完
	a = 2
	b = 3
	c = 0

	// 回调函数
	return func() (int64, error) {
		d := a + b + c
		return d, nil
	}
}

type Clazz struct {
	Name    string
	Id      *int64
	Student *string
	Cz      *clazz
}

func TestOperatingNewStruct() {
	var a = new(clazz)
	var b *clazz = &clazz{}
	a.Name = "123"
	fmt.Println(a)

	b.Name = "123"
	fmt.Println(b)

	var c = new(Clazz)
	c.Name = "123"
	c.Id = new(int64)
	*c.Id = 123
	c.Student = new(string)
	*c.Student = "123"
	c.Cz = new(clazz)
	c.Cz.Name = "123"

	fmt.Println(c)
}

func TestXErrorsTryCatch() {
	var err error
	err = xerrors.New("123")
	fmt.Println(err)

	err = xerrors.Errorf(": %w", err)

	if errors.Is(err, xerrors.New("123")) {
		fmt.Println("捕获!")
	}
}

func TestChannelGetStarted() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	fmt.Printf("len:%v, cap:%v chan:%v\n", len(ch), cap(ch), ch)

	// ch <- 30
}

func putNum(intChan chan<- int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func TestCalculatePrimeUsingChannel() {
	// 1. 将要求的数字放入 channel
	ch := make(chan int, 1000)
	go putNum(ch)

	// 2. 开启 4 个协程，从 channel 中取出数字，判断是否为素数，如果是，放入另一个 channel
	primeChan := make(chan int, 3000)

	exitChan := make(chan bool, 4)

	for i := 0; i < 4; i++ {
		go func() {
			for {
				num, ok := <-ch
				if !ok {
					// 说明 intChan 取完了
					break
				}

				var flag bool = true
				for i := 2; i < num; i++ {
					if num%i == 0 {
						flag = false
						break
					}
				}

				if flag {
					primeChan <- num
				}
			}

			exitChan <- true
		}()
	}

	// 3. 等待所有的协程完成工作
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		close(primeChan)
	}()

	for {
		val, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%v\n", val)
	}

}

func TestChannelInsteadOfErrorGroup() {
	// n := 0
	// for i := 0; i < 100; i++ {
	// 	n += i
	// }

	eg := errgroup.Group{}
	eg.SetLimit(10)
	var n = 0
	// var theLock sync.Mutex
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(func() error {
			n += i
			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

func TestUmarshalStructSlice() {

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	jsonData := []byte(`[
		{
			"name": "Alice",
			"age": 30
		},
		{
			"name": "Bob",
			"age": 25
		},
		{
			"name": "Charlie",
			"age": 22
		}
	]`)

	var people []*Person

	err := json.Unmarshal(jsonData, &people)
	if err != nil {
		log.Fatal(err)
	}
	r, _ := json.Marshal(people)
	fmt.Println(string(r))
}

func GenerateSql() {
	sql := ``

	tables := []string{
		"dm_data_ecom.sim_update_inner_and_outer_migration", "dm_data_ecom.temp_", "dm_data_ecom.ods_prd_price_compare_sla_goods_pool_with_properties_1d", "dm_data_ecom.ods_price_compare_feature_completed_hf", "dm_data_ecom.ods_price_compare_feature_retrigger_cal_hf", "dm_data_ecom.sync_src_diff_to_sku_price_variety", "dm_data_ecom.sku_price_variety_abase_dump", "dm_data_ecom.price_compare_sla_task_summary", "dm_data_ecom.price_compare_sla_price_update", "dm_data_ecom.price_compare_sla_same_update", "dm_data_ecom.price_compare_sla_sku_goods_pool", "dm_data_ecom.ecom_mmc_apply_center_diff_repair_sku_light_weight", "dm_data_ecom.ecom_mmc_apply_center_diff_repair_spu_light_weight", "dm_data_ecom.create_price_cmp_task_sla", "dm_data_ecom.dump_ecom_product_price_cmp_gather_msg", "dm_data_ecom.price_compare_online_task_sku_re_calculate_debounce_hour", "dm_data_ecom.price_compare_online_task_sku_re_calculate_debounce", "dm_data_ecom.ods_src_realtime_product_price_power_hour", "dm_data_ecom.price_compare_caiji_same_sku_i_d_aidp", "dm_data_ecom.inner_same_shop_30_d_price_change", "dm_data_ecom.diff_between_lowest_sku_and_platform", "dm_data_ecom.dwd_thunder_task_sku_public", "dm_data_ecom.sim_update_inner_and_outer", "dm_data_ecom.cases_have_been_deduplicated_sku_ids", "dm_data_ecom.outer_sim_sku_lowest_update_time", "dm_data_ecom.hash_json_diff_inner", "dm_data_ecom.business_sku_stock_diff", "dm_data_ecom.activity_dim_diff_consistency", "dm_data_ecom.activity_thunder_task_sku_diff", "dm_data_ecom.cmp_aidp_mark_data_set", "dm_data_ecom.cmp_aidp_gather_data_set", "dm_data_ecom.activity_dataset_619", "dm_data_ecom.commop_price_control_center_dataset", "dm_data_ecom.cmp_offline_mistake_feedback_data_set", "dm_data_ecom.cmp_offline_data_set", "dm_data_ecom.cmp_offline_data_mid_table", "dm_data_ecom.dataset_consistency_of_cmp_diff", "dm_data_ecom.equivalent_abase_diff", "dm_data_ecom.dwd_thunder_task_sku", "dm_data_ecom.dwd_ecom_product_price_cmp_res_msg", "dm_data_ecom.ods_ecom_product_price_cmp_res_msg", "dm_data_ecom.bytegraph_pack_diff", "dm_data_ecom.ods_bg_price_compare_sku_node", "dm_data_ecom.ods_data_edc_gr_product_change_rmq", "dm_data_ecom.hash_json_diff", "dm_data_ecom.same_skus_day_by_platform", "dm_data_ecom.thunder_sku_diff", "dm_data_ecom.ods_goods_price_compare_task_config",
	}
	for _, table := range tables {
		sql += `select count(*) 
			from ` + table + `
			where date= '${date}' union
			`
	}

	fmt.Print(sql)
}

func IpV6Compress() {
	ipv6Addr := "0000:0000:0000:0000:0000:0000:0000:0001"
	ip := net.ParseIP(ipv6Addr)
	if ip == nil {
		fmt.Println("Invalid IPv6 address")
		return
	}
	ipv6Compressed := ip.String()
	fmt.Println(ipv6Compressed)
}

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func Reverse(head, tail *LinkedList) (*LinkedList, *LinkedList) {
	var prev = tail.Next
	var p = head
	for prev != tail {
		var nextNode = p.Next
		p.Next = prev
		prev = p
		p = nextNode
	}
	return tail, head
}

func ReverseLinkedListGroupByK(l *LinkedList, k int) *LinkedList {
	var h = &LinkedList{} // head Node
	h.Next = l
	var pre = h

	for l != nil {
		// find [a, b] while b - a + 1 = k
		var tail = pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return h.Next
			}
		}

		// reverse by [a, b]
		var nextNode = tail.Next
		l, tail = Reverse(l, tail)
		pre.Next = l
		tail.Next = nextNode
		pre = tail
		l = tail.Next
	}

	return h.Next

}

func SortSliceHavingOneElement() {
	type S struct {
		Value string
	}
	var a []*S

	sort.Slice(a, func(i, j int) bool {
		return a[i].Value < a[j].Value
	})

	fmt.Printf("A: %v\n", ToJson(a))

}

func XerrorsTest() {
	f1 := func() error {
		return fmt.Errorf("error in f1")

	}

	f2 := func() error {
		err := f1()
		return xerrors.Errorf("error in f2: %w", err)
	}

	f3 := func() error {
		err := f2()
		return xerrors.Errorf("error in f3: %w", err)
	}

	err := f3()
	if err != nil {
		fmt.Printf("INFO: %+v\n", err)
	}

}

func Dial() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write([]byte("ZHUIOJZHSIDKHAJKDHAKJHGHKJHAGKJ"))

	conn.Close()
}

func Listen() {
	l, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening.............")

	c, err := l.Accept()

	fmt.Println("Accept.............")

	c.Close()
	buf := make([]byte, 1024)
	c.Read(buf)
	fmt.Printf("INFO: %s\n", string(buf))
}

func SortMapValues() {
	m := map[string][]string{
		"1": {"2", "3", "1"},
	}

	for _, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
	}

	fmt.Println(ToJson(m))
}

func DualThreadPrint100() {
	f := func(i int) {
		fmt.Println(i)
	}

	aTob := make(chan int)
	bToa := make(chan int)
	done := make(chan int)
	go func() {
		if len(done) > 0 {
			return
		}

		for i := 1; i <= 100; i += 2 {
			f(i)
			aTob <- i + 1
			<-bToa
		}
	}()

	go func() {
		for {
			i := <-aTob
			f(i)
			if i == 100 {
				done <- 1
			} else {
				bToa <- i
			}
		}
	}()

	time.Sleep(time.Second * 3)
}

func LinkedListReverse(head *LinkedList) *LinkedList {
	p := head.Next

	// 头结点
	var result *LinkedList = &LinkedList{}

	for p != nil {
		s := &LinkedList{
			Val:  p.Val,
			Next: nil,
		}
		// 头插入
		s.Next = result.Next
		result.Next = s

		p = p.Next
	}

	return result
}

func PrintLinkedList(h *LinkedList) {
	p := h.Next
	for p != nil {
		fmt.Printf("%v ", p.Val)
		p = p.Next
	}

	fmt.Println()
}

func BitCalculate() {
	var a = 0x01010131

	var b = 1<<12 - 1

	var c = a & b

	fmt.Printf("%012b\n", c)
}

func TestnanFloat() {
	a := "XXX"
	x := cast.ToFloat64(a)
	fmt.Println(x) // NaN????
	fmt.Printf("%.2f", x)
}

func main() {

	// test2(true)
	// test3()
	// test4()
	// test5()
	// test7()
	// test8()
	// test9()
	// test12()
	// test15()
	// test18()
	// test19()
	// test22()
	// test23()
	// test24()
	// x := test25()
	// if x == nil {
	// 	print("NULL")
	// }
	// test26()
	// test27()
	// test28()
	// test29()
	// test30()
	// test31()
	// test32()
	// test33()
	// test34()
	// test35()
	// test36()
	// test37()
	// test38()
	// test_between_string_and_fmtSprint()
	// test_nil_struct_invoke_its_method()
	// test_context_reference_address()
	// test_sync_once()
	// test_map_nil_access()
	// test_time_format()
	// test_int_array_return()
	// test_slice_index()
	// TestMarshalSequencely()
	// TestTwoDimsArraySlice()
	// TestClosureUse()
	// TestOperatingNewStruct()
	// TestXErrorsTryCatch()
	// TestChannelGetStarted()
	// TestCalculatePrimeUsingChannel()
	// TestChannelInsteadOfErrorGroup()
	// TestUmarshalStructSlice()
	// test15()
	// GenerateSql()
	// IpV6Compress()
	// SortSliceHavingOneElement()
	// XerrorsTest()
	// SortMapValues()
	// DualThreadPrint100()
	// h := &LinkedList{}
	// r := h
	// for i := 1; i <= 3; i++ {
	// 	s := &LinkedList{
	// 		Val:  i,
	// 		Next: nil,
	// 	}

	// 	r.Next = s
	// 	r = s
	// }
	// PrintLinkedList(h)
	// res := LinkedListReverse(h)
	// PrintLinkedList(res)
	// Dial()
	// Listen()

	// BitCalculate()
	TestnanFloat()
}
