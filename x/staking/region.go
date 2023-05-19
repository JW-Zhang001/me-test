package staking

import (
	"fmt"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"
	"me-test/config"
	"me-test/tools"
	"strings"
)

func (k *Keeper) NewRegion(privKey, regionId, name, validator string) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, _ := tools.GetAccAddress(privKey)
	fromAddr := fromAccAddr.String()
	zap.S().Info("NewRegion/fromAddr: ", fromAddr)
	zap.S().Info("NewRegion/regionId: ", regionId)
	zap.S().Info("NewRegion/name: ", name)

	msg := stakepb.NewMsgNewRegion(fromAddr, regionId, name, validator)
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("ValidateBasic error")
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, privKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *Keeper) RandRegionKey() []string {
	regionKey := make([]string, 0, 256)
	regionName = strings.Replace(regionName, "\n", "", -1)
	regionNameSli := strings.Split(regionName, ",")
	for _, rn := range regionNameSli {
		name := strings.Split(rn, "-")
		regionKey = append(regionKey, name[0])
	}
	return regionKey
}

var regionName = `ABW-阿鲁巴,
AFG-阿富汗,
AGO-安哥拉,
AIA-安圭拉,
ALA-奥兰,
ALB-阿尔巴尼亚,
AND-安道尔,
ARE-阿联酋,
ARG-阿根廷,
ARM-亚美尼亚,
ASM-美属萨摩亚,
ATA-南极洲,
ATF-法属南部和南极领地,
ATG-安提瓜和巴布达,
AUS-澳大利亚,
AUT-奥地利,
AZE-阿塞拜疆,
BDI-布隆迪,
BEL-比利时,
BEN-贝宁,
BES-荷兰加勒比区,
BFA-布基纳法索,
BGD-孟加拉国,
BGR-保加利亚,
BHR-巴林,
BHS-巴哈马,
BIH-波黑,
BLM-圣巴泰勒米,
BLR-白俄罗斯,
BLZ-伯利兹,
BMU-百慕大,
BOL-玻利维亚,
BRA-巴西,
BRB-巴巴多斯,
BRN-文莱,
BTN-不丹,
BVT-布韦岛,
BWA-博茨瓦纳,
CAF-中非,
CAN-加拿大,
CCK-科科斯群岛,
CHE-瑞士,
CHL-智利,
CHN-中国,
CIV-科特迪瓦,
CMR-喀麦隆,
COD-刚果民主共和国,
COG-刚果共和国,
COK-库克群岛,
COL-哥伦比亚,
COM-科摩罗,
CPV-佛得角,
CRI-哥斯达黎加,
CUB-古巴,
CUW-库拉索,
CXR-圣诞岛,
CYM-开曼群岛,
CYP-塞浦路斯,
CZE-捷克,
DEU-德国,
DJI-吉布提,
DMA-多米尼克,
DNK-丹麦,
DOM-多米尼加,
DZA-阿尔及利亚,
ECU-厄瓜多尔,
EGY-埃及,
ERI-厄立特里亚,
ESH-西撒哈拉,
ESP-西班牙,
EST-爱沙尼亚,
ETH-埃塞俄比亚,
FIN-芬兰,
FJI-斐济,
FLK-福克兰群岛,
FRA-法国,
FRO-法罗群岛,
FSM-密克罗尼西亚联邦,
GAB-加蓬,
GBR-英国,
GEO-格鲁吉亚,
GGY-根西,
GHA-加纳,
GIB-直布罗陀,
GIN-几内亚,
GLP-瓜德罗普,
GMB-冈比亚,
GNB-几内亚比绍,
GNQ-赤道几内亚,
GRC-希腊,
GRD-格林纳达,
GRL-格陵兰,
GTM-危地马拉,
GUF-法属圭亚那,
GUM-关岛,
GUY-圭亚那,
HMD-赫德岛和麦克唐纳群岛,
HND-洪都拉斯,
HRV-克罗地亚,
HTI-海地,
HUN-匈牙利,
IDN-印度尼西亚,
IMN-马恩岛,
IND-印度,
IOT-英属印度洋领地,
IRL-爱尔兰,
IRN-伊朗,
IRQ-伊拉克,
ISL-冰岛,
ISR-以色列,
ITA-意大利,
JAM-牙买加,
JEY-泽西,
JOR-约旦,
JPN-日本,
KAZ-哈萨克斯坦,
KEN-肯尼亚,
KGZ-吉尔吉斯斯坦,
KHM-柬埔寨,
KIR-基里巴斯,
KNA-圣基茨和尼维斯,
KOR-韩国,
KWT-科威特,
LAO-老挝,
LBN-黎巴嫩,
LBR-利比里亚,
LBY-利比亚,
LCA-圣卢西亚,
LIE-列支敦士登,
LKA-斯里兰卡,
LSO-莱索托,
LTU-立陶宛,
LUX-卢森堡,
LVA-拉脱维亚,
MAF-法属圣马丁,
MAR-摩洛哥,
MCO-摩纳哥,
MDA-摩尔多瓦,
MDG-马达加斯加,
MDV-马尔代夫,
MEX-墨西哥,
MHL-马绍尔群岛,
MKD-北马其顿,
MLI-马里,
MLT-马耳他,
MMR-缅甸,
MNE-黑山,
MNG-蒙古,
MNP-北马里亚纳群岛,
MOZ-莫桑比克,
MRT-毛里塔尼亚,
MSR-蒙特塞拉特,
MTQ-马提尼克,
MUS-毛里求斯,
MWI-马拉维,
MYS-马来西亚,
MYT-马约特,
NAM-纳米比亚,
NCL-新喀里多尼亚,
NER-尼日尔,
NFK-诺福克岛,
NGA-尼日利亚,
NIC-尼加拉瓜,
NIU-纽埃,
NLD-荷兰,
NOR-挪威,
NPL-尼泊尔,
NRU-瑙鲁,
NZL-新西兰,
OMN-阿曼,
PAK-巴基斯坦,
PAN-巴拿马,
PCN-皮特凯恩群岛,
PER-秘鲁,
PHL-菲律宾,
PLW-帕劳,
PNG-巴布亚新几内亚,
POL-波兰,
PRI-波多黎各,
PRK-朝鲜,
PRT-葡萄牙,
PRY-巴拉圭,
PSE-巴勒斯坦,
PYF-法属波利尼西亚,
QAT-卡塔尔,
REU-留尼汪,
ROU-罗马尼亚,
RUS-俄罗斯,
RWA-卢旺达,
SAU-沙特阿拉伯,
SDN-苏丹,
SEN-塞内加尔,
SGP-新加坡,
SGS-南乔治亚和南桑威奇群岛,
SHN-圣赫勒拿阿森松和特里斯坦-达库尼亚,
SJM-斯瓦尔巴和扬马延,
SLB-所罗门群岛,
SLE-塞拉利昂,
SLV-萨尔瓦多,
SMR-圣马力诺,
SOM-索马里,
SPM-圣皮埃尔和密克隆,
SRB-塞尔维亚,
SSD-南苏丹,
STP-圣多美和普林西比,
SUR-苏里南,
SVK-斯洛伐克,
SVN-斯洛文尼亚,
SWE-瑞典,
SWZ-斯威士兰,
SXM-荷属圣马丁,
SYC-塞舌尔,
SYR-叙利亚,
TCA-特克斯和凯科斯群岛,
TCD-乍得,
TGO-多哥,
THA-泰国,
TJK-塔吉克斯坦,
TKL-托克劳,
TKM-土库曼斯坦,
TLS-东帝汶,
TON-汤加,
TTO-特立尼达和多巴哥,
TUN-突尼斯,
TUR-土耳其,
TUV-图瓦卢,
TZA-坦桑尼亚,
UGA-乌干达,
UKR-乌克兰,
UMI-美国本土外小岛屿,
URY-乌拉圭,
USA-美国,
UZB-乌兹别克斯坦,
VAT-梵蒂冈,
VCT-圣文森特和格林纳丁斯,
VEN-委内瑞拉,
VGB-英属维尔京群岛,
VIR-美属维尔京群岛,
VNM-越南,
VUT-瓦努阿图,
WLF-瓦利斯和富图纳,
WSM-萨摩亚,
YEM-也门,
ZAF-南非,
ZMB-赞比亚,
ZWE-津巴布韦`
