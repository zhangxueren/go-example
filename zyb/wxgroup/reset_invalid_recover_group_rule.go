package wxgroup

import (
	"fmt"
	"io"
	"net/http"
)

type RequestParams struct {
	WxGroupId   string `json:"wxGroupId"`
	TeamGroupId int    `json:"teamGroupId"`
	Classify1Id int    `json:"classify1Id"`
	Classify2Id int    `json:"classify2Id"`
	LabelId     int    `json:"labelId"`
	GradeId     int    `json:"gradeId"`
	UrlTemplate string `json:"-"`
	Headers     map[string]string
	Cookies     map[string]string
}

// 刷新年级错配导致督学群回收失败任务规则，重新接回
func ResetInvalidRecoverGroupRule() {

	// 下周一diff出已成功继承数据，修改群分组
	wxGroupIds := []string{"10799676622116429", "10714768838666395", "10899231656907801", "10863314898129292", "10857506315571947", "10951635520556525", "10970787398618093", "10823593351228861", "10718003703515122", "10858450484421195", "10763457050609083", "10788724686815631", "10909028538952125", "10860320759731124", "10739412886481709", "10865475240067361", "10901994938719279", "10734126218747692", "10912743034358620", "10814016455423563", "10811247512190806", "10955664142868100", "10793312668510779", "10893769862975704", "10890023877068083", "10798028377183048", "10757140748416842", "10785329610551358", "10903567993674831", "10722048995124919", "10911964459057830", "10905560310576576", "10733949415603651", "10782465345164887", "10853676490516653", "10950488881188021", "10743959566737257", "10861932565447411", "10747310633543449", "10799986370897024", "10922964732196045", "10970201446662753", "10940946057514926", "10707657203631333", "10877865812269182", "10730070246720485", "10884032014874235", "10745967542777152", "10928648394537745", "10701203823442773", "10857191362380243", "10850792787015635", "10880789763996654", "10864009184772628", "10853633278164627", "10820833490817268", "10740734484149133", "10707457514489101", "10911313186492212", "10938138340102932", "10781986427492257", "10907930674217961", "10973526359014966", "10820691588362661", "10711871286492593", "10912319352639351", "10741410383450670", "10964953340327231", "10727925674322386", "10939343646625024", "10901378008435752", "10910653924710825", "10724939351927994", "10796486357052245", "10879501434331124", "10929087868911932", "10923259227446789", "10815653331197381", "10777808658827456", "10880510714037166", "10752891376215400", "10790627103725058", "10731656811250989", "10943213606968615", "10745554082795384", "10916014365351156", "10783122549214942", "10840187420806006", "10946252628216890", "10937732839099597", "10774021975618061", "10941090371396241", "10964542297864798", "10947339717751372", "10757264837714342", "10859372362384400", "10871587617564145", "10716510723050905", "10807818096220386", "10905934645362358", "10974545157032306", "10785503646678065", "10744298374919179", "10976385986156912", "10723637080205241", "10854316499034160", "10716178212223676", "10808401707849768", "10886845248315053", "10745396859757479", "10815049351450783", "10757825204851570", "10839615309430916", "10760506190417334", "10896081639097167", "10743549811840007", "10921261427735998", "10831103207710451", "10868137383302876", "10907959850628534", "10937034485517065", "10940302237266996", "10833393046164252", "10773314938198111", "10813828695072720", "10893911230713245", "10931237175148504", "10789614460353643", "10813659459465651", "10802192387619356", "10738547295746071", "10867472073836156", "10882663934470632", "10735164620454571", "10958636718888944", "10770814093740717", "10765144528861529", "10912304366077321", "10960907522806925", "10763887321137488", "10735288892463902", "10871954481552841", "10761879936625330", "10726526453232798", "10788973616428570", "10740472342265308", "10758768992007611", "10842659081542008", "10943879127306449", "10705396308541410", "10952508239280003", "10794597287571625", "10859869760305145", "10903885706524907", "10890453844434340", "10860610482640266", "10937631837287894", "10788251603777613", "10845875457849157", "10804181842397302", "10916953258765316", "10777631697898948", "10713311198029461", "10865593998329621", "10789479017828603", "10971112713130801", "10819179022554982", "10896348039066566", "10956850569525817", "10726755360040414", "10882931451708132", "10874793294005953", "10706273502922038", "10812923415487296", "10885456680733463", "10892238893602520", "10784798273134946", "10803203050724001", "10841330369668276", "10888761731608603", "10816390121077808", "10949694629148271", "10756882630976563", "10752498672233498", "10827900349235944", "10880515610559628", "10696348225294102", "10931396735503876", "10703340267198027", "10715306131996542", "10911903436785943", "10707688683091962", "10864641677437047", "10907320781057556", "10933448359778653", "10765065636420608", "10908058824911732", "10753966486473683", "10816832841582667", "10738948776010581", "10859023759339754", "10708573975090070", "10751295083782777", "10789982949040455", "10921111647594052", "10728998639521948", "10936974259989924", "10891280479351126", "10944474220154741", "10773851525190183", "10757775203743423", "10797456333053245", "10717029743674571", "10773402191000300", "10741984594780427", "10759621227465693", "10848733779953138", "10814413841702944", "10808605376823325", "10926616000075769", "10817407225403514", "10807250316788917", "10911484516676997", "10707956023811610", "10803512652189694", "10844666742893281", "10976463080655988", "10864917517958571", "10977497044719098", "10757217569558665", "10751173219525111", "10759953939407670", "10934406751043862", "10747868060897616", "10799117879378629", "10848453399006448", "10919970068950521", "10708712100235272", "10803883188079920", "10844504702698845", "10818642817558013", "10948669140831539", "10899394685889776", "10881096275191054", "10739178193823590", "10919664183867586", "10878806705271891", "10909575782541432", "10852051975616488", "10804912611154567", "10903893018309222", "10934057355633542", "10952302583240198", "10923840498340931", "10849674764193205", "10827156240637608", "10917919163331946", "10930584889677972", "10826272645881192", "10714556772959466", "10769662373499781", "10778404657921532", "10831716550753829", "10955145828050445", "10827453908645321", "10836603202322388", "10964878151552491", "10760690030151272", "10761157025785077", "10893743373950480", "10795854944627565", "10828073375127532", "10761725465222774", "10833580756616172", "10938334582786820", "10844826535494928", "10764868227686791", "10921270446530574", "10808051569427206", "10761603975768393", "10817674714396046", "10697357822782419", "10864904716915079", "10899666624644185", "10814990671308276", "10893516560505240", "10859184524334299", "10828795914782069", "10961500387725338", "10787514752196911", "10870882179518140", "10769536405056808", "10856664023872140", "10749578765341058", "10818523416191310", "10850468238191132", "10719230168650673", "10716463908920819", "10859712182853175", "10828126948901184", "10792917555953089", "10798699184440207", "10915656999378021", "10935368219267308", "10698423089291562", "10926278573629776", "10850906261952842", "10892436534810518", "10855288821091339", "10746147590293648", "10836282464003482", "10949855523379110", "10913057776166832", "10773068434246541", "10829080642165648", "10758920516448898", "10917751918412663", "10764832858496339", "10778653438314729", "10900285609444638", "10956743879140843", "10969702324233485", "10847414350418312", "10809399876052284", "10881355530689857", "10898769106967436", "10899094014869674", "10951011663947485", "10869372716052356", "10806784660276773", "10930201276217160", "10769006039550401", "10743302702537850", "10923590138714338", "10820116690220568", "10702522258286821", "10925378152891913", "10923841547470431", "10942453809072292", "10961408024169982", "10916441703033872", "10933360278767655", "10780976808601907", "10930339199182312", "10890695037698465", "10876079911797787", "10866249582198453", "10950884386845542", "10915341021216922", "10756168087288577", "10913305362218036", "10735714371829958", "10837352219319487", "10967218543877648", "10921997120422918", "10891763251118567", "10780575073248924", "10876220432788765", "10974452313491106", "10862194207752464", "10952995039068561", "10875372992522140", "10914669066377057", "10870062120545923", "10760854465871547", "10731409753658486", "10707056981604163", "10874329368570388", "10743925552291805", "10892435975311544", "10738553522525100", "10931067111912753", "10760156450613979", "10840126806163220", "10751947444641174", "10964591477718099", "10948615154989059", "10889774460682752", "10735835995850800", "10865197206547797", "10877014000968149", "10811745710951125", "10797463179499534", "10806182262572227", "10806964448846163", "10869277604692632", "10907859056231422", "10712606130914288", "10849997430712691", "10760640123775425", "10780343708402167", "10707462704135149", "10961284640522293", "10926756022505850", "10737026007661081", "10907117770247950", "10773181863837733", "10890355843490355", "10918717537803541", "10738359293955182", "10801772532589603", "10788103912117461", "10795093582627298", "10802034680541730", "10835856799069947", "10812642036196774", "10919878002295071", "10956012754771135", "10816567111954606", "10699507601980301", "10908317344026821", "10755598085224769", "10891696061608574", "10744692182841638", "10918299537506219", "10856253880957745", "10719783973331823", "10839722424999684", "10804692107856237", "10834240242907904", "10879423678207612", "10969065928309226", "10941372744147365", "10884193659619752", "10833787214424817", "10715714971308165", "10965646064765457", "10950332645923964", "10847003728271516", "10839179860097932", "10899154047463943", "10754707082416559", "10925702739058310", "10729500261952537", "10738097079809704", "10747033200178399", "10933542786545473", "10726522932018936", "10940278427717237", "10929468674664509", "10788510672945386", "10711747152657892", "10772571181027870", "10795015849976523", "10943313393664557", "10857836038559797", "10733360067656980", "10882468661361805", "10873529502299242", "10877329426754479", "10857395950944278", "10729696504905965", "10822554975765200", "10771790412852550", "10941810019537714", "10798979712114126", "10731479281828590", "10827483855659118", "10784697138007995", "10706251237508851", "10952810166702555", "10905059709048457", "10742120585853993", "10887259190621029", "10841559919005795", "10821077364672685", "10798505020949706", "10733441018113707", "10786895387015678", "10700392431800322", "10954693890072239", "10715315280124539", "10907399737817403", "10828487486927979", "10876585210824048", "10894888244220924", "10940204875627615", "10880683299233694", "10974716235019039", "10971087687427377", "10908377599078614", "10954377858786444", "10748944931561859", "10831743122677142", "10866486904655378", "10748392297573697", "10810738653818657", "10806012066431968", "10872208707093483", "10725952133011735", "10891933946176648", "10916474752529914", "10723890980111017", "10758242814940912", "10937345120550043", "10945121946334705", "10943399188798781", "10773816037570077", "10797994947862666", "10758856124332773", "10801841855615658", "10759522016106921", "10724286104715287", "10873584357005974", "10716908269365143", "10731455330424011", "10723151044923799", "10751292129761259", "10869786355840894", "10785713133218875", "10915817024533308", "10730815248418711", "10855768000263973", "10945877129111454", "10953471129324927", "10908855969659728", "10891443744195102", "10752538509676968", "10725910222033715", "10827124460761096", "10737593905342878", "10975461094014643", "10922866192819517", "10892275541457314", "10909635229449110", "10855427912504011", "10827795888111206", "10909493451630231", "10889787856378788", "10804632210622412", "10710361874253061", "10790828034056650", "10953989609966594", "10950731545845465", "10744795273773167", "10840503220905754", "10806443948703572", "10826486562329521", "10816278633962949", "10800082693220286", "10833395670517151", "10866989223137897", "10908181957250599", "10898953850551723", "10810900025606102", "10909147668270698", "10940995416171010", "10754981587310770", "10924878249974946", "10886051326179875", "10822307777505975", "10776335804262842", "10819570247695907", "10896362087601998", "10939364351957823", "10968085798608875", "10793391200788117", "10722535583035912", "10957107553065517", "10892741197624341", "10921795810980961", "10892158437360707", "10775686191133947", "10738579941687105", "10703773661917004", "10919723839400328", "10906298293743640", "10780413246623063", "10889335275749520", "10865742333554366", "10803028732054276", "10928345644099452", "10731418975123086", "10957792523522482", "10810703706103388", "10818734292321275", "10905920777827813", "10814776089025496", "10851154906592003", "10872069644928137", "10896090476006664", "10778694429125595", "10868462640720713", "10748798389162596", "10755418893196771", "10883797038262576", "10713000617821581", "10935126284608769", "10784742077169183", "10784100129671767", "10856396716977151", "10942588439374187", "10728996804141822", "10757089306349381", "10814886451711468", "10717671896019679", "10763200338354146", "10768406300714723", "10723379993064513", "10901102670138124", "10915998651268784", "10802735293224682", "10849151137851073", "10697558917746535", "10782907435396447", "10826836289805259", "10835976974651622", "10852145328953950", "10900543372052501", "10806475086074319", "10722150918003690", "10697203996059774", "10746600019865468", "10755560340600653", "10880903533535428", "10762163194957937", "10812573660479678", "10710127214384561", "10719513341256644", "10730756972634895", "10947832993543094", "10942896501751169", "10733975652254868", "10786537788589263", "10891030411527951", "10742452534567712", "10824608214212672", "10880103427466989", "10800856244378405", "10838515978062289", "10753789450367623", "10874292737906877", "10825336030844240", "10728567722135655", "10812962515879052", "10911775668421573", "10941670728763371", "10872861301191581", "10787046813420554", "10704842231476609", "10923588724699743", "10911399877747046", "10920772602866445", "10771956332536645", "10923331941062793", "10812461555726638", "10715182245755509", "10715504674255131", "10953460181164587", "10891946217828604", "10868169926249223", "10802919446117546", "10950134047288938", "10935879100560044", "10970684544385657", "10773608888068677", "10917249416759848", "10892435451102439", "10771109552713608", "10921538354047910", "10730387503197302", "10969395439914406", "10780511691825452", "10752584804874352", "10735001659213265", "10795866583931533", "10773467509371397", "10965111016983539", "10740676198591663", "10864599551126528", "10944134938780287", "10697417475648741", "10707873940804388", "10830729146355317", "10810824090847654", "10748328273255758", "10799193064610792", "10760652545062859", "10939844249611963", "10833908044174574", "10779420671121934", "10953353705172186", "10783300195755281", "10710341781801563", "10959449483623765", "10764227951160843", "10873301834954785", "10930585783422477", "10740008161973990", "10851026090226303", "10842501154630630", "10933778882716175", "10757144098268251", "10937285789596189", "10791908088443958", "10841728678247106", "10881008962132647", "10833481528719227", "10849045234419140", "10841141401325496", "10752512568607243", "10898799051877004", "10752588605714539", "10897146524076532", "10867324460098330", "10743313717038601", "10800619427250513", "10882475703348016", "10879648904307801", "10719432644283303", "10931697390244408", "10758458728740675", "10972190738876226", "10701501140385938", "10866203641926145", "10912682089722048", "10815309730677717", "10871436733062748", "10914531363792106", "10769952905593271", "10788934956742996", "10718759938702632", "10895408007300120", "10719736164110588", "10822178610565418", "10916348717463303", "10826856976797848", "10878534422193111", "10901337771372564", "10765831881268167", "10952387140023142", "10893741790573534", "10781226652076339", "10884976795016727", "10856445942906686", "10841244866155547", "10931084326703918", "10852188118121090", "10749654415692359"}

	for _, wxGroupId := range wxGroupIds {
		err := sendRequest(wxGroupId)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func sendRequest(wxGroupId string) error {
	urlStr := fmt.Sprintf("https://wxtools.zuoyebang.cc/wxqk-go/toolsv2/flushrecovergrouprule?wxGroupId=%s&teamGroupId=172&classify1Id=9163&classify2Id=9176&labelId=31&gradeId=5", wxGroupId)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", "RANGERS_WEB_ID=991ac774-eed9-4805-9fbb-fe9ea61621c8; RANGERS_SAMPLE=0.04972195991847683; ZYBIPSCAS=IPS_243c284344cd8e86af69527a9f5738211721133774; Hm_lvt_c33960c712441eec1b994580263ccb1a=1719804603,1721738711; Hm_lpvt_c33960c712441eec1b994580263ccb1a=1721738711; HMACCOUNT=19A5FA5D6642C585; uid=zhangxueren")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Sec-Ch-Ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("WxGroupId:", wxGroupId, " Response Status:", resp.Status, " Response Body:", string(body))

	return nil
}
