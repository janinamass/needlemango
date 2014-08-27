package needlemango

import (
	"strconv"
	"strings"
	"io/ioutil"
)

/*
 SUBSTITUTION MATRIX
*/
type SubstitutionMatrix struct{
	name string
	data map[string]int
}

func MakeSubstitutionMatrix(name string) SubstitutionMatrix {
    return SubstitutionMatrix{name, map[string]int{"**":1}}
}


func (sm SubstitutionMatrix) GetName() string {
	return sm.name
}

func (sm SubstitutionMatrix) GetVal(keyA string, keyB string) int {
	if val,ok := sm.data[keyA+keyB]; ok {
    	return val
	}else if val,ok := sm.data[keyB+keyA]; ok{
		return val
	}
	//not in map
	panic("Aaah! Unexpected key.")
}

func (sm *SubstitutionMatrix) SetMap(mapName string){
	if mapName == "EBLOSUM62"{
		sm.data = nil
		sm.data = map[string]int{"**":1,"*A":-4,"*C":-4,"*B":-4,"*E":-4,"*D":-4,"*G":-4,"*F":-4,"*I":-4,
"*H":-4,"*K":-4,"*M":-4,"*L":-4,"*N":-4,"*Q":-4,"*P":-4,"*S":-4,
"*R":-4,"*T":-4,"*W":-4,"*V":-4,"*Y":-4,"*X":-4,"*Z":-4,"A*":-4,
"AA":4,"AC":0,"AB":-2,"AE":-1,"AD":-2,"AG":0,"AF":-2,"AI":-1,
"AH":-2,"AK":-1,"AM":-1,"AL":-1,"AN":-2,"AQ":-1,"AP":-1,"AS":1,
"AR":-1,"AT":0,"AW":-3,"AV":0,"AY":-2,"AX":0,"AZ":-1,"C*":-4,
"CA":0,"CC":9,"CB":-3,"CE":-4,"CD":-3,"CG":-3,"CF":-2,"CI":-1,
"CH":-3,"CK":-3,"CM":-1,"CL":-1,"CN":-3,"CQ":-3,"CP":-3,"CS":-1,
"CR":-3,"CT":-1,"CW":-2,"CV":-1,"CY":-2,"CX":-2,"CZ":-3,"B*":-4,
"BA":-2,"BC":-3,"BB":4,"BE":1,"BD":4,"BG":-1,"BF":-3,"BI":-3,
"BH":0,"BK":0,"BM":-3,"BL":-4,"BN":3,"BQ":0,"BP":-2,"BS":0,
"BR":-1,"BT":-1,"BW":-4,"BV":-3,"BY":-3,"BX":-1,"BZ":1,"E*":-4,
"EA":-1,"EC":-4,"EB":1,"EE":5,"ED":2,"EG":-2,"EF":-3,"EI":-3,
"EH":0,"EK":1,"EM":-2,"EL":-3,"EN":0,"EQ":2,"EP":-1,"ES":0,
"ER":0,"ET":-1,"EW":-3,"EV":-2,"EY":-2,"EX":-1,"EZ":4,"D*":-4,
"DA":-2,"DC":-3,"DB":4,"DE":2,"DD":6,"DG":-1,"DF":-3,"DI":-3,
"DH":-1,"DK":-1,"DM":-3,"DL":-4,"DN":1,"DQ":0,"DP":-1,"DS":0,
"DR":-2,"DT":-1,"DW":-4,"DV":-3,"DY":-3,"DX":-1,"DZ":1,"G*":-4,
"GA":0,"GC":-3,"GB":-1,"GE":-2,"GD":-1,"GG":6,"GF":-3,"GI":-4,
"GH":-2,"GK":-2,"GM":-3,"GL":-4,"GN":0,"GQ":-2,"GP":-2,"GS":0,
"GR":-2,"GT":-2,"GW":-2,"GV":-3,"GY":-3,"GX":-1,"GZ":-2,"F*":-4,
"FA":-2,"FC":-2,"FB":-3,"FE":-3,"FD":-3,"FG":-3,"FF":6,"FI":0,
"FH":-1,"FK":-3,"FM":0,"FL":0,"FN":-3,"FQ":-3,"FP":-4,"FS":-2,
"FR":-3,"FT":-2,"FW":1,"FV":-1,"FY":3,"FX":-1,"FZ":-3,"I*":-4,
"IA":-1,"IC":-1,"IB":-3,"IE":-3,"ID":-3,"IG":-4,"IF":0,"II":4,
"IH":-3,"IK":-3,"IM":1,"IL":2,"IN":-3,"IQ":-3,"IP":-3,"IS":-2,
"IR":-3,"IT":-1,"IW":-3,"IV":3,"IY":-1,"IX":-1,"IZ":-3,"H*":-4,
"HA":-2,"HC":-3,"HB":0,"HE":0,"HD":-1,"HG":-2,"HF":-1,"HI":-3,
"HH":8,"HK":-1,"HM":-2,"HL":-3,"HN":1,"HQ":0,"HP":-2,"HS":-1,
"HR":0,"HT":-2,"HW":-2,"HV":-3,"HY":2,"HX":-1,"HZ":0,"K*":-4,
"KA":-1,"KC":-3,"KB":0,"KE":1,"KD":-1,"KG":-2,"KF":-3,"KI":-3,
"KH":-1,"KK":5,"KM":-1,"KL":-2,"KN":0,"KQ":1,"KP":-1,"KS":0,
"KR":2,"KT":-1,"KW":-3,"KV":-2,"KY":-2,"KX":-1,"KZ":1,"M*":-4,
"MA":-1,"MC":-1,"MB":-3,"ME":-2,"MD":-3,"MG":-3,"MF":0,"MI":1,
"MH":-2,"MK":-1,"MM":5,"ML":2,"MN":-2,"MQ":0,"MP":-2,"MS":-1,
"MR":-1,"MT":-1,"MW":-1,"MV":1,"MY":-1,"MX":-1,"MZ":-1,"L*":-4,
"LA":-1,"LC":-1,"LB":-4,"LE":-3,"LD":-4,"LG":-4,"LF":0,"LI":2,
"LH":-3,"LK":-2,"LM":2,"LL":4,"LN":-3,"LQ":-2,"LP":-3,"LS":-2,
"LR":-2,"LT":-1,"LW":-2,"LV":1,"LY":-1,"LX":-1,"LZ":-3,"N*":-4,
"NA":-2,"NC":-3,"NB":3,"NE":0,"ND":1,"NG":0,"NF":-3,"NI":-3,
"NH":1,"NK":0,"NM":-2,"NL":-3,"NN":6,"NQ":0,"NP":-2,"NS":1,
"NR":0,"NT":0,"NW":-4,"NV":-3,"NY":-2,"NX":-1,"NZ":0,"Q*":-4,
"QA":-1,"QC":-3,"QB":0,"QE":2,"QD":0,"QG":-2,"QF":-3,"QI":-3,
"QH":0,"QK":1,"QM":0,"QL":-2,"QN":0,"QQ":5,"QP":-1,"QS":0,
"QR":1,"QT":-1,"QW":-2,"QV":-2,"QY":-1,"QX":-1,"QZ":3,"P*":-4,
"PA":-1,"PC":-3,"PB":-2,"PE":-1,"PD":-1,"PG":-2,"PF":-4,"PI":-3,
"PH":-2,"PK":-1,"PM":-2,"PL":-3,"PN":-2,"PQ":-1,"PP":7,"PS":-1,
"PR":-2,"PT":-1,"PW":-4,"PV":-2,"PY":-3,"PX":-2,"PZ":-1,"S*":-4,
"SA":1,"SC":-1,"SB":0,"SE":0,"SD":0,"SG":0,"SF":-2,"SI":-2,
"SH":-1,"SK":0,"SM":-1,"SL":-2,"SN":1,"SQ":0,"SP":-1,"SS":4,
"SR":-1,"ST":1,"SW":-3,"SV":-2,"SY":-2,"SX":0,"SZ":0,"R*":-4,
"RA":-1,"RC":-3,"RB":-1,"RE":0,"RD":-2,"RG":-2,"RF":-3,"RI":-3,
"RH":0,"RK":2,"RM":-1,"RL":-2,"RN":0,"RQ":1,"RP":-2,"RS":-1,
"RR":5,"RT":-1,"RW":-3,"RV":-3,"RY":-2,"RX":-1,"RZ":0,"T*":-4,
"TA":0,"TC":-1,"TB":-1,"TE":-1,"TD":-1,"TG":-2,"TF":-2,"TI":-1,
"TH":-2,"TK":-1,"TM":-1,"TL":-1,"TN":0,"TQ":-1,"TP":-1,"TS":1,
"TR":-1,"TT":5,"TW":-2,"TV":0,"TY":-2,"TX":0,"TZ":-1,"W*":-4,
"WA":-3,"WC":-2,"WB":-4,"WE":-3,"WD":-4,"WG":-2,"WF":1,"WI":-3,
"WH":-2,"WK":-3,"WM":-1,"WL":-2,"WN":-4,"WQ":-2,"WP":-4,"WS":-3,
"WR":-3,"WT":-2,"WW":11,"WV":-3,"WY":2,"WX":-2,"WZ":-3,"V*":-4,
"VA":0,"VC":-1,"VB":-3,"VE":-2,"VD":-3,"VG":-3,"VF":-1,"VI":3,
"VH":-3,"VK":-2,"VM":1,"VL":1,"VN":-3,"VQ":-2,"VP":-2,"VS":-2,
"VR":-3,"VT":0,"VW":-3,"VV":4,"VY":-1,"VX":-1,"VZ":-2,"Y*":-4,
"YA":-2,"YC":-2,"YB":-3,"YE":-2,"YD":-3,"YG":-3,"YF":3,"YI":-1,
"YH":2,"YK":-2,"YM":-1,"YL":-1,"YN":-2,"YQ":-1,"YP":-3,"YS":-2,
"YR":-2,"YT":-2,"YW":2,"YV":-1,"YY":7,"YX":-1,"YZ":-2,"X*":-4,
"XA":0,"XC":-2,"XB":-1,"XE":-1,"XD":-1,"XG":-1,"XF":-1,"XI":-1,
"XH":-1,"XK":-1,"XM":-1,"XL":-1,"XN":-1,"XQ":-1,"XP":-2,"XS":0,
"XR":-1,"XT":0,"XW":-2,"XV":-1,"XY":-1,"XX":-1,"XZ":-1,"Z*":-4,
"ZA":-1,"ZC":-3,"ZB":1,"ZE":4,"ZD":1,"ZG":-2,"ZF":-3,"ZI":-3,
"ZH":0,"ZK":1,"ZM":-1,"ZL":-3,"ZN":0,"ZQ":3,"ZP":-1,"ZS":0,
"ZR":0,"ZT":-1,"ZW":-3,"ZV":-2,"ZY":-2,"ZX":-1,"ZZ":4}
	}else{}
}

/*
SEQUENCE
*/
type Sequence struct{
	header string
	sequence string
}
func MakeSequence(header string, sequence string) Sequence{
	return Sequence{header:header, sequence:sequence}
}
func (s Sequence) GetHeader() string{
	return s.header
}

func (s Sequence) GetSequence() string{
	return s.sequence
}

func toFasta(seq Sequence) string{
	var res string
	res = seq.header+"\n"+seq.sequence
	return res
}

/*
SCORE
*/
type Score struct {
	seqA Sequence
	seqB Sequence
	sm SubstitutionMatrix
	res float64
}
func (s *Score ) setScore(newVal float64){
	s.res=newVal
}
func PrettyPrint(score Score) string{
	var res string
	res = score.seqA.header+"\t"+score.seqB.header+"\t"+strconv.FormatFloat(score.res,'f', 2, 32)+"\n"
	return res
}
/*
NMW
*/
func Nmw(seqA Sequence, seqB Sequence, substMat SubstitutionMatrix )Score{
	//twoD array to iterate over
	m := len(seqA.sequence)+1
	n := len(seqB.sequence)+1
	var s Score
	s = Score{seqA:seqA, seqB:seqB, sm: substMat, res:0}
	//2d Scoring Matrix
	var mat = make([][]float64, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]float64, n)
	}
	// Scoring matrix done

	//Boolean matrix to allow for gap extension
	var isGap = make([][]bool, m)
	for i := 0; i < m; i++ {
		isGap[i] = make([]bool, n)
	}
	var seqA_arr []string
	seqA_arr = strings.Split(seqA.sequence,"")

	var seqB_arr []string
	seqB_arr = strings.Split(seqB.sequence,"")

	var gapOpenPenalty float64
	gapOpenPenalty = -10

	var gapExtendPenalty float64
	gapExtendPenalty = -0.5

	var left float64
	var top float64
	var diag float64

	for i:=1; i<=len(seqA_arr);i++{
		v := seqA_arr[i-1]
		for  j:= 1;j <= len(seqB_arr);j++ {
			isGap[i][j] = false

			if isGap [i][j - 1] {
				left = mat[i][j-1]+gapExtendPenalty
			}else {left = mat[i][j-1]+gapOpenPenalty}

			if isGap [i - 1][j] {
				top = mat[i-1][j]+gapExtendPenalty
			}else {top = mat[i-1][j]+gapOpenPenalty}

			w := seqB_arr[j - 1]
			diag = mat[i-1][j-1]+float64(s.sm.GetVal(v, w))
			mat[i][j] = maxOfThree(top, left, diag)

			if mat[i][j] == top || mat[i][j] == left {
				isGap[i][j] = true
			}

		}
	}
	s.setScore(getMax(mat))
	return s
}


func getMax(twod [][]float64) float64{
	var max float64
	var n int
	var m int
	n = len(twod)
	m = len(twod[n-1])
	max = -1
	//last col max
	lastCol := n-1
	lastRow := twod[lastCol]
	for _,v :=range lastRow{
		if max < v{
			max = v
		}
	}
	//last row max
	lastRowIndex := m-1
	for i:=0;i<n-1;i++{
		if twod[i][lastRowIndex]>max{
			max = twod[i][lastRowIndex]
		}
	}
	return(max)
}


func maxOfThree(x float64, y float64, z float64)float64{
	if x >= y && x >= z{
		return x
	}else if y >= x && y >= z{
		return(y)
	}
	return z
}


type FastaReader struct{
	file string
}

func(fr FastaReader) getSequences()[]Sequence{
	var res []Sequence
	dat, err := ioutil.ReadFile(fr.file)
    check(err)
	var seqarr []string
	seqarr = strings.Split(string(dat),">") // now everything should start with a header
	for _,v := range(seqarr){
		var tmp []string
		tmp = strings.Split(v,"\n")
		shortHeader := strings.Split(tmp[0]," ")[0]
		res = append(res, Sequence{header:shortHeader, sequence:strings.Join(tmp[1:],"")})
	}
	return res
}

type FastaReaderStd struct {
	inputString string
}

func(frs FastaReaderStd) getSequences()[]Sequence {
	var res []Sequence
	var seqarr []string
	seqarr = strings.Split(string(frs.inputString), ">") // now everything should start with a header
	for _, v := range (seqarr) {
		var tmp []string
		tmp = strings.Split(v, "\n")
		shortHeader := strings.Split(tmp[0], " ")[0]
		res = append(res, Sequence{header:shortHeader, sequence:strings.Join(tmp[1:], "")})
	}
	return(res)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
