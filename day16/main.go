package main

import "fmt"

var m = map[byte]byte{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
}
var res int

func main() {
	//str := "A052E04CFD9DC0249694F0A11EA2044E200E9266766AB004A525F86FFCDF4B25DFC401A20043A11C61838600FC678D51B8C0198910EA1200010B3EEA40246C974EF003331006619C26844200D414859049402D9CDA64BDEF3C4E623331FBCCA3E4DFBBFC79E4004DE96FC3B1EC6DE4298D5A1C8F98E45266745B382040191D0034539682F4E5A0B527FEB018029277C88E0039937D8ACCC6256092004165D36586CC013A008625A2D7394A5B1DE16C0E3004A8035200043220C5B838200EC4B8E315A6CEE6F3C3B9FFB8100994200CC59837108401989D056280803F1EA3C41130047003530004323DC3C860200EC4182F1CA7E452C01744A0A4FF6BBAE6A533BFCD1967A26E20124BE1920A4A6A613315511007A4A32BE9AE6B5CAD19E56BA1430053803341007E24C168A6200D46384318A6AAC8401907003EF2F7D70265EFAE04CCAB3801727C9EC94802AF92F493A8012D9EABB48BA3805D1B65756559231917B93A4B4B46009C91F600481254AF67A845BA56610400414E3090055525E849BE8010397439746400BC255EE5362136F72B4A4A7B721004A510A7370CCB37C2BA0010D3038600BE802937A429BD20C90CCC564EC40144E80213E2B3E2F3D9D6DB0803F2B005A731DC6C524A16B5F1C1D98EE006339009AB401AB0803108A12C2A00043A134228AB2DBDA00801EC061B080180057A88016404DA201206A00638014E0049801EC0309800AC20025B20080C600710058A60070003080006A4F566244012C4B204A83CB234C2244120080E6562446669025CD4802DA9A45F004658527FFEC720906008C996700397319DD7710596674004BE6A161283B09C802B0D00463AC9563C2B969F0E080182972E982F9718200D2E637DB16600341292D6D8A7F496800FD490BCDC68B33976A872E008C5F9DFD566490A14"
	str := "9C0141080250320F1802104A08"
	n := len(str)
	bits := make([]byte, n*4)
	for i := 0; i < n; i++ {
		num := m[str[i]]
		for j := 0; j < 4; j++ {
			bits[i*4+j] = (num >> (3 - j)) & 1
		}
	}

	_, num := dfs(0, bits)
	fmt.Println(num)
	fmt.Println(res)
}

func dfs(index int, bits []byte) (int, int) {
	if index >= len(bits) {
		return index, 0
	}

	res += parseNum(index, index+3, bits)
	index += 3
	typ := parseNum(index, index+3, bits)
	index += 3

	num := 0
	if typ == 4 { // literal value
		for bits[index] == 1 {
			num <<= 4
			num |= parseNum(index+1, index+5, bits)
			index += 5
		}
		num <<= 4
		num |= parseNum(index+1, index+5, bits)
		index += 5
	} else { // operator
		lengthType := bits[index]
		index++

		var subNum int
		var nums []int
		if lengthType == 1 {
			subPacketNum := parseNum(index, index+11, bits)
			index += 11
			for i := 0; i < subPacketNum; i++ {
				index, subNum = dfs(index, bits)
				nums = append(nums, subNum)
			}
		} else {
			subPacketLen := parseNum(index, index+15, bits)
			index += 15
			tmp := index
			for index < tmp+subPacketLen {
				index, subNum = dfs(index, bits)
				nums = append(nums, subNum)
			}
		}

		switch typ {
		case 0: // sum
			for i := 0; i < len(nums); i++ {
				num += nums[i]
			}
		case 1: // product
			num = nums[0]
			for i := 1; i < len(nums); i++ {
				num *= nums[i]
			}
		case 2: // minimum
			num = nums[0]
			for i := 1; i < len(nums); i++ {
				if nums[i] < num {
					num = nums[i]
				}
			}
		case 3: // maximum
			num = nums[0]
			for i := 1; i < len(nums); i++ {
				if nums[i] > num {
					num = nums[i]
				}
			}
		case 5: // greater than
			if nums[0] > nums[1] {
				num = 1
			}
		case 6: // less than
			if nums[0] < nums[1] {
				num = 1
			}
		case 7: // equal
			if nums[0] == nums[1] {
				num = 1
			}
		}
	}

	return index, num
}

func parseNum(start, end int, bits []byte) int {
	num := 0
	for start < end {
		num <<= 1
		num |= int(bits[start])
		start++
	}

	return num
}
