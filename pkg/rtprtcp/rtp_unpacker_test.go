// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtprtcp

import (
	"encoding/hex"
	"testing"

	"github.com/q191201771/naza/pkg/bele"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/naza/pkg/assert"
)

// #85
func TestAvcCase1(t *testing.T) {
	// single sps
	// single pps
	// FUA IDR
	ss := []string{
		"a06000013778b64c1921d51867640032ad84010c20086100430802184010c200843b5014005ad370101014000003000400000300ca100002",
		"806000023778b64c1921d51868ee3cb0",
		"806000033778b64c1921d5187c85b8000006bff0ffee6021caad8ffdd001c15e27306e7fc6bf36ca8b6bc1411afef158dc64d75094f5b2bbddec364e6904cc9a14b4069bb6b6cb6d6ab8f132e77e3324291351b52e4a58fce30dd7b64313d208df50ab636423abce5a0dcc6c3ff8397c84250042244aafef705063debcdd7fe6c6c2fd41b4c6251fe4ca318c40e4bf2eb373246f14ad1623f9f5356154b02c0c8f53fcd8b6ad452da457b48ba704ec227f16b59d9e7c0d50423bb7f78059a68d26aef86bc94ebb27723c6eee018b45977028b400e474d40f7464fcb5f0f292b23a2324486bdd500fe72b115bd93bdf5f0f9207daedfc02e4b8bed2c2b8a4b081b00c3a9172e1ce34222a870d06b925a9263c59276d4ad6bd82bdab07f8e01b3ae3263899948d345f83e4ff127b3c4d2721d3543880a0ea72d9b13ea0a4dd64ca4617f3d63b0182f9cc37c917dc620853f12487d608bc76fe3d5bc039fc9df161d4d6ca2cae273f7893ffbe78ba9e0e85217c9c2b6214e48e831cf4eed74f1e70f3050fcd72dd466d1a3d00f500b884c141625f5daba37c20e4623387406a6930819f08c0e3321ea7695d92e8da38e5d926767926ee5ebce20e4eecbd7aea2b8756489271ba7c2977c4e568a4a25c82b159cc3f5f6575f1b96ad1c589267947258bc1e62d5eed1a17a99219c019fd9167754310ad3c2ef1acbf304fe0b5342d09ae20b4b4fc49f5231ca744977d3a73ce2821f8cc76f1c0149ae8d5dadf44bf87ad989a311ce57ad19e4579bc8cdc309eb28954e875441f9c135182eca8f3016aa5792ec3d1926da40694465dfd625b8c8c836792f4cee619bcce39a71c8055827cbd5c22fedfcb7626534d07385df9247b6338e9bc68c1ee8903a4fc9309decc5ce2ca60a7331ccd15f285ddb9be2134173fada2150ce68dddc4d2ef397624d74cb0ff9dab86664e17a0bc76ec86875736e68133e3bcdf5419cd6ce19ff1108bbe7e39ea36578ea8e8a4cbe6c850095467f22c1c220aa8203f8ee16d39de0f414db0f46dbc2f7f34b4dc7062aeed316c991522efb7fd8ac10d510596638f0bc330005317ab085787db0b26ed4c6086513d059ceef9274d6788a9bf30abb1cf6b353d9e57103e10a299df9d77c688cc0d69af84f319e97f55db30d419fdf9a028b4185fa1cae5f2fdfadcb0abc68ca03968a8b4801d31d8096451b1dd54d159d81b3cf8554a43a65a80fc1581eb00719de7c606201c31af22d0fef3208889c4f0a5c0ea06ecdb13e6d99cff25f449d1fe90c3d7beadc52595b12d9e58fc647ebbc30ac4131fad7dfad3a6ea9d51948b095d9aafd48a0545378c2083406a2248f4876e104d76d22ce0a552848d89d603392dad6486554a58dd8fb154b70e52d4c4e1f093d73d731563bebf9fe61493779ba791df3dc65430cf27d00292609f6c09a682e9f4b03a4e3507459ae06ff07e66fa992b54f8d0837cf806d9921ac916b1fe064adb76e7cedfa4e2c5a9126a0f7a7927e5686b990cc87f89612b983cb81783e0763a548648ef73d855a0afeaf78192047b060fcf8b1016a46016de75983d7e57c70d5fa5259012e465531f953b27ca67b554620a5e194386b8fd6d784e989d151b7326d028ff28b2707d376147ca3657a60daa7158e630b9def41c579a02b24c6b5cbe5161c985bf5a6ba41b033be092c4a590098aca5a4a6b8aef024863ef403fe4cb4dc2010dedd3cc1c5f7c6a123e2d7f8fd8feef793afb2402a763df927316b337d83808f13354f4706af395556caa6ec0efb724a7abe587b8eff333d64623ad7556a735dcd04bfe88c8f57327625aef8f25a0b722499de91a509ad383650d3ba250e16c5dd2671d67039b9404c174863a9af5c2738dd293770488c034c1f4f9c0d1cfb8f02c9bcecfb24a17ad06c7163788f22c8bbeb30b26423ad22515aca1916a28f716aad8b970623f054537448d74ee4822",
	}

	testHelperTemplete(t, base.AvPacketPtAvc, 90000, 128, ss, func(rtpPackets []RtpPacket) []base.AvPacket {
		return []base.AvPacket{
			{
				Timestamp:   10340642,
				PayloadType: base.AvPacketPtAvc,
				Payload:     testHelperAddPrefixLength(rtpPackets[0].Raw[12:]),
			},
			{
				Timestamp:   10340642,
				PayloadType: base.AvPacketPtAvc,
				Payload:     testHelperAddPrefixLength(rtpPackets[1].Raw[12:]),
			},
		}
	})
}

func TestHevcCase1(t *testing.T) {
	// single vps
	// single sps
	// single pps
	// single sei
	ss := []string{
		"a060d3a38a27999a01c0125940010c01ffff016000000300b0000003000003007bac0901",
		"a060d3a48a27999a01c01259420101016000000300b0000003000003007ba003c08010e58dae4914bf37010101008001",
		"a060d3a58a27999a01c012594401c0f2c68d03b240000003",
		"a060d3a68a27999a01c012594e01e504ebc3000080000003",
	}

	testHelperTemplete(t, base.AvPacketPtHevc, 90000, 128, ss, func(rtpPackets []RtpPacket) []base.AvPacket {
		return []base.AvPacket{
			{
				Timestamp:   25753900,
				PayloadType: base.AvPacketPtHevc,
				Payload:     testHelperAddPrefixLength(rtpPackets[0].Raw[12:]),
			},
			{
				Timestamp:   25753900,
				PayloadType: base.AvPacketPtHevc,
				Payload:     testHelperAddPrefixLength(rtpPackets[1].Raw[12:]),
			},
			{
				Timestamp:   25753900,
				PayloadType: base.AvPacketPtHevc,
				Payload:     testHelperAddPrefixLength(rtpPackets[2].Raw[12:]),
			},
			{
				Timestamp:   25753900,
				PayloadType: base.AvPacketPtHevc,
				Payload:     testHelperAddPrefixLength(rtpPackets[3].Raw[12:]),
			},
		}
	})
}

func TestHevcCase2(t *testing.T) {
	// vps40,sps42,pps44,sei4e
	// 85 = 2 + 24 + 2 + 35 + 2 + 8 + 2 + 10
	ss := []string{
		"80601191bd9ee38884bc42f46001001840010c01ffff016000000300b0000003000003007bac09000023420101016000000300b0000003000003007ba003c08010e58dae4932f4dc040404020000084401c0f2f03c9000000a4e01e5040c7500008000",
	}
	out, _ := hex.DecodeString("0000001840010c01ffff016000000300b0000003000003007bac090000000023420101016000000300b0000003000003007ba003c08010e58dae4932f4dc0404040200000000084401c0f2f03c90000000000a4e01e5040c7500008000")

	testHelperTemplete(t, base.AvPacketPtHevc, 90000, 128, ss, func(rtpPackets []RtpPacket) []base.AvPacket {
		return []base.AvPacket{
			{
				Timestamp:   35347852,
				PayloadType: base.AvPacketPtHevc,
				Payload:     out,
			},
		}
	})
}

func TestAacCase1(t *testing.T) {
	ss := []string{
		"80e10e9a56843e4cf0bdf2fe00102d10214e6c425f74815f92415f94415f924100000114008a004027f313d564a770026fd01203c9cbac420e1aecaa41efb4619391309daa7938b4b905989c5c293c024819e4234eb324934327e3f083c98c2220137bc9b5b2bae5809351710e4bfec0dec11fffe0ab6e568f8357fb74c3461e892fc3a4e22ac512fa73bdc13004b9d73e21c5221387922692939994e38ab1516ea2cfb64df8cffb5468a81e4704f255a5139f7f21d622b944e0b08221916d4ba6e10a31b2a148c38642d1c8b2b12a15681c0f0b589160feb32b43c7ebade413590916a934148ca7915d6250904e5172a3ecc612ab23fd4431b0eef591be52155e4a0ac8564108012689c4df89221892e50d8fca64a747bb2313908211004222ea6a848132e87935978e03845905ccc057c743c7424e64e2338bfe52090d4a5a961f35ec5e14befbc2b3d41f89bdfe949fc2dd40e141b13e397f84f7e1ef43f303df77407bdacc01c2832c3d3df65fe3bfde0001f37f85f4a021f27baf11e4f388fd5ce464337648fc89e6043a0c8268a44f0e024e112910c8b9bc282ee2e5587762b2752c08a422ac99558f904273f858d9de9938ef2502493ae3217c247410eb39395f084a117f6486161103c5cab0881d7c24ffd7123cdcac226bc292c2dc213d04f0b26dd864e6e149c3a4468a48d86124209e172166c5b331e434c8cae4276e191acc2682d692a51210c1ce2254102e048435919442449f2c57bb271364fbb67f1d79364ee004a28c88e75dd53210eef613186531f22908a5253a87e66423fe1bffef7dd95fed588acf815908808fefff2ffef581fe23b2b8e9fdfbd7e3f7071c05d40ec89401efbdf87beee9a9c1eff2d054c0f7ff5febc514514514514514514517d7fe0faf145145111fc0dd8218d9f9d064ee8895cb84ee848064e7593530495f87ffa276e76549a4e6c0e381270c24ea288e0e1d8ad270f0dc243210e413a84b160e085b10a416be12011022b31354b202020771fe376bf6afe387e67fe0fee8087bee102bf415fa0afd0400000e3e3e3fe001f8f8ff53eff60fafab04fe0877121cebad11c16048271174dc251cbdd3da84e0c9ba23dd312b26500020154a269d9c4520270444131665313103f0f7412d40de5c59cfff0f6383440fc588e986f10b35ddbfdb6be780f30eb25f49fd7e479062097adbc044ba3e0c956ef043c31d188a23d453e83ca71ca44df64805f4514916e71c7e760674293094980642012d104ba2cacd260456e2e93a8c1f492739f9dd445aba9241088de157cb3d976e20888a41b34808f3a4ca181528c989641e0a187dc5443494c896e4421711580a84276a12a229662108289464d105b1513f80841836f009c987f86969641eea943c86c475685fb09043ae947ac5d22ff6ff7c1d380a48885c2a7d03905668ec9fd1fcdf7fc700450094d16f034c90003df9240b851764658e101efb8ea83855071d6871c5126a8385556aa0db094e571d9dc7271352b846810c4e24000000000000038e60384c4f09800000021e17eb447ca381218ff036078221ec1e7843a466acc013c0130585c7518330993235855b71158610843210aa9ba05c22ea5984de5e1227e3e4d2dd942a1659f5320232b45bb6c6418df87b313c727fc3f1d170a6ecb0b74e4449209804a3bf270c98518182b79dd74412c24f7e4f45131a889442502df3928282006d0f165c6e4d39059e7404eb1092044c67267817686dc7dbe82281f1d2f65ddaf70e67250c77adbe39d09be32b8f0007d2f905140d3c34fb6784075ce77075dfb069da091017344c80b6b9a36ff1d93c22889d5a7ccaa2073ccf278e1f8e9c38e1c9a8f502b8e9c84938542ae3a7e1313c26240000000000000384c4f1cc4f09890000021f49fc3847d1b76e8613938122e759d2494b9e4634fc812ab747e3132beb075bf23fed67a0943a245082648a4d06e11166a25f110a24c0c993964042ff0a8d24e4108c489514499641221f84c013c06008e0301c2e0038e6009e0327c7326000001c",
		"80e10e9b5684424cf0bdf2fe00102ce8217c4445fd4080000000000021e10fe2c047f365f8a423e48fb4c479f78f48f9e7c2647b178548f27f1693ecdec1b1f704b84989705ea8436bb5241b849b83253444f41189a621383b227b2aa4db0721db25369113cd2598864316d24dd1f1cb7c72df0b049c182460e0f8ee4b846115622efa44950683513c2cee38dcecd270f05c7544d4a274d44b2f90272ed9352894b51182c2583a846952e12092a6a2106913834eef824e0d42188a6455049a95c72590a6a232367c71404a9ab843f1c9a4df84e10e40f6838e5a2534e47187e12d71cd07093486821f1d7f1dc6123e3c83ed71c8c43198fe387207390b56b8e49a8ae13c8c6254e490934c8cdc6903e0091f00429d699b5041ef20f3f0a7218c3933e3c9da8dc25a231ad10c65a24fd011895489f0e4ee61c8d5cd1124d25630842ae1ff449e1c84f55608c489c2552106c10b812798918226672119f548e1cf334f99e2139b109c1391cc26a58c46f6289cc2108d9da924112f01ba2bd4f2081495314996e13246278f51291188032e431a4e12e13c7cc211884e428996e12228948915260896731b5314848c511246b16458923844939f9ae10f82a664aa46aa49e3c845292512a90cc61c844b84f0f548444ccb565725d50c89af92363889a7dd5267e912adf22f2127d4217e1910c32724d74867ea24de12318e4a6cd231ad71c3923c620739239c8cd3f0b1b8ecd2734e4669c81ce4ec1c81e3118d378edc2520e4621085c21021ea1493244267904c8424490151c499a254310912a913c82210920bc8049c71e4030f1841e102210f08f9543c2867f081008480424c2193d7ffde103fff7ffdffdc7e1c7e19cf081e103ffce10201fc80019543fc9c2a88843c2872a07f97f6003f93f943f3801fedfc9fb1fecf7f7e00ffb003cffdffec1ff7fd8ffb7e0ffee16717feffeff8200ffb07e7ffedfeffb201fc9ff6ff600000ff60f7fc705ef7fdbfdfdffbc0ff7ff60f7dff700f7ffec1ff7e3a2000000000008bbfa81000000000bd757353a04bb17e6748fdcefa6647ea8fd4223f3a7ab114f2227c33784f63025af4510f7f4e27773a4fe09f4e238fe2e470f9b2649a45f8827c9b9f12d5d5267b533da22439374822fb39d9e4819920b270842306411832099e31139c89e690d963f84a24423b7211309c905f2b1b8e908df2129a8e381201090861b12d906848c80916c8267a247315bb2b85392861acce44f0c81e1103d62038640272610f08d250c24715778e078edb210c39dd5c71e4a3ceceeace8524a571c2e3d7fff091544148e11e41181e3b1f84a79d23118e5e115c77098f15ffe90296704d6323125cbd0e4c864df533b2b3a28843d11280b211e3e3f8f9d55ffdcef1b852902ab8e978e14947449e1cef009c8804433b2a3c81eed121e1287a2928278370812718e46462498564a844cac823403f762741126225c7e3c8c44312b65135baefa5d244e5288c59b812884a55d9149c55138b23c389560d06b217f17fc6e12413ad4ab7a84961271646570925ac8de5607309c55105bbc3082e8115d0e12e13a13bff64e8442214110c4bbeb1189124c0900c093a41380193a2100a32a1484a4117b31f55c7af212a5112ac830bc8d284320c464141394fb3e511293a17ff849efe3a5c7caffd7fffffe63d781353f1fbf1ebfffec7f5bff5c75fc76463f5f1fdd923fd24d11394b3aaff5f959612e224d44bc8cac1ceebfff7ccf3a86739d1fc757e7c1ffecd73a2fff78f561ffe7ff84ce799e271c270872242112138e1091c848e4e3b13843f087e3a4e10e01c2c100027065f0b0785091972f85964a0cbe120f1d8201c2c1e1208108d1803848dc2cbe3b2f848dc7308425c600e39180384c271d97c725871d97c768f1c964a5d1e3b2c3204be3b2f85964e5d100e392c09e0301c7301c2e038e6240385c4f1cc07098000000e13001c72c71c6f096384c4f0b89e396385b01c2e2784c4871db1c771201c7359c270b0e3bc2f8ed700000000000e",
		"80610e9c5684464cf0bdf2fe00102db8211a0455fd2003dbd44f3d9b002dd3e3cdef4097d20fbaa47f5b9c188fa9524fb075d2184b84f1fb320cdf93e3be3427dcee120512392e584b07048931a4f0e923073642e6dec4391c3e62ecd4133d32184ab32c421124114e9384e09c72c64fb1c2d1a8cf33442756a91261c963544a49c9cd9ba0909a7256ad119b3493dfc2e089435109782234d5c2d8278db0047619e20d81c2d3233b124605badade62233a5909d4c9406e769c4675b25025811916f8e378ed1e3ad216f177e21528909ef2640907d1201a846fdf22679008654bfc75fc2849bde40c4206792904226be44350811f3f9a7e7f1c092bab228ba4db0c84d41044d24dc571d0929b74922d909a19f833e4f24f791c24327a37926bc9bda45a426d9f2ac424186490de1024aedbe3b2c966aa708de16ef1cc713b61e3b348a0241eac7f0ff29c767135371e4e250dff84cef7484cba467cfe15e429c122b90420e0c848210590848a646e1091241090b210689182423228912138e12b1a240212418e4a1c3270864470adca84027210a8901d12279b6fc2255a112850894e7da6ea2c04b2d2091084a14eceb6485d21005d2565e49e4236615a0f26eba41e427225e3e9dc2bc9d9b645f88232e16777e3c7fffe45ef2407f1dbb9d41ffa24e793b242207f1d3f084a8884e4248d89b50108c488462134026725472094139139f8e412473ccd367f9f534fe127900c3230c248269549c7084c24261864e1be57864a1865786442195a191b91bb289c18c46e412583b96edcf7fc241e16e5bd83ac20f1c8dc76080138d18841824ae4621060d610495cc1dbf078e4125060f1d824a0c12706e0567078ec1e12e7097385e2bc773dc7727c733dc264f8e64c0009a327c2e78038ee7b8ee4c967b27c271fe39cd78e71f0000e178ff09c7f8e735000e178f871ce6bc2f1fe178f8709c7c384c971ce3fc274e0e3bc7f8ef1f00e39c7f84e9c04375bae3bac0e3bbae13580001c73a170bace1359c26eb85dd13dd6b385d6709c7f84e6bc271f0233b9a870bc7c384e3fc738f871ce6bc739a8709cd4236fe903aefae6bf1df8bd4e7384ec025f57e7cc012f89fdf88fe2c7e6490f055c27f0f37c46748215b5a4373009f60f7d10f1c4823c131043ddfdf09632313562f3b6148c6cd11e0d9120dcee74804e25f24190467064d449ab271365c2e1b8e6009e9701c25827560da5548448849b85c79778fc82d7c2a0943093420843864e6cd270e590c12b8ec3274d5c2349e2da4a8d1e16d920c9e15f2728849864570c8ad72941949177be5995292c9be0f1d8c46d4de38702318e466bf858846e41995244e4a8e09025cbb2acc97081cf95e4e9dc29ed8237152cd5bb5646ab320c3e0644b56592ffd89369117378ea0901c45c0e382b324130378475980b3cb8f4ad4421225b912d4194e0ddb071f412358b9091c23b2048e16a9124925021f1d270ad24493ff6224059ee2250f6a91808e3af221935253ba0b7440bac64a449b67852f082270961762eec844162bb246565905c2e393ff609ae7feb2ec3e543fe73214c248454306648248f0ffbbda57641cecf990f32a2d924f2ca8424d21332ff5a4a33ee93ddb24836492396ef2913c62442cfb26ed2d8aba99fc296c63d125c1635d65a90bdae44a524d2f1d9374b7f1a7c2ddc722d2fecffbd9e520879122ecd2b4f427085fbf122a2ee83fade3907f5c49652280930d0e103951776af2b44cec8ffeffab2bae5e5925a2ee5925a3f5a4d68ad9a4674e24a6d6eaa0e9d053b8e4ee115b6109d4c8aa771d6f1d500106b4841836ec621060918372dcc191b9078e4125060f1d824a0c1aca081283738ec1f7bef785171dc4f1cc071cc4f0b80e3b8000025b3c070b8900e3b89e39802789c0709bae3bc2f8e6e8000385dd709bae3bc2c00385dd071de17c26b384d6070bac0e13a871dd6709c2c38e6b38e6b0038eeb384e1604375bae39ac0e3bbae17580001c71fc26b385d6709bae17744b75ace1359c2eb384dd70bac023d0b741c26b0385d671dd6071cdd",
		"80e10e9d5684464cf0bdf2fe00102db871cdd0709ba0e0",
	}

	testHelperTemplete(t, base.AvPacketPtAac, 48000, 128, ss, func(rtpPackets []RtpPacket) []base.AvPacket {
		return []base.AvPacket{
			{
				Timestamp:   30239734,
				PayloadType: base.AvPacketPtAac,
				Payload:     rtpPackets[0].Raw[12+4:],
			},
			{
				Timestamp:   30239756,
				PayloadType: base.AvPacketPtAac,
				Payload:     rtpPackets[1].Raw[12+4:],
			},
			{
				Timestamp:   30239777,
				PayloadType: base.AvPacketPtAac,
				Payload:     append(rtpPackets[2].Raw[12+4:], rtpPackets[3].Raw[12+4:]...),
			},
		}
	})
}

func TestAacCase2(t *testing.T) {
	ss := []string{
		"80e104558424aa36b06db689003000c000300030de04004c61766335372e3130372e31303000422008c11838211004608c1c212004608c1c",
		"80e104568424b636b06db68900102d7020824ecd7ec9037f01bf9cdfe1af65fe2f8929dfff5754fd6be1a938fffbd772fdae7b717ad7cfff15d712f87c756ebec000025c1677100007f1aec2801869302bacc0003ecf821400b0a24b8fca00008c628be329f92424a4989200000001598ba8883e3d8a700000000c0dca33e12a4380000000000064f04e8520b013044271d16820000000000000834fdd38181d45ee0000000000000aec04ce29d0a41a9fe09352c00000000000000025a110233005e4f613687cf3000e8d0000fec7dcb47a7d482a131082024d2655350e6204812e36826cb0cd6c45100950191294826310282650139f33fcf5345951945c1ba584e00f071139b28839e4081a1c44d5389b49d704d4e26b06547d0b0c8162ecaad8b44d0cbf3a5020b7d4a69925d6d0c82265a8dc7376c2a9aa58af99d73d8e504b4796c377322153aaee260019f4b2a20900fe61db59003fe622264ca9f39f6d6c36ed9fc1ca84bbcb451e3b9789f5dfc5ba65d0cb0dda74503aa48962118439b889152a16595d624caa02214e0c5a9012b1e8326760a222c5114bec74e7f9fd3af736db38fa3f977887413632f7634bc5ad4f93a058c6fff3316573d9f2c834128c5caa7268749c22063642935a8a57371f5aa2c991ae925664268630db8b26621094cff8fd94bee0267844e6a4810741aad616577d4116df07ad91124998dfae2288248e720105dc8b1c84a356938445f66d3054a7273536e12cf5d88ebae01228ae945da7a0a1107109c29ddc84173c80da4254ca0436fbf1ea6ee1d669a21ddf64ca8c1459d4142d420e45aa62514feb5352c0722f16f04f1ab25183bafa0c9803dc44081f48a4350ede00bac2031ebbe1f2e05632091e060aca20d2dbc91864de8b3e21186722c310801ad884ce722f836fd2ad4372e750c991b2fd23338e7e791953890e1e5707a41121884d29054a210e0904088a679148884151284d9bc9497105970429181170237ad67655484244875ac02320a4cd66c79b900b52c624b011203dcfb8a7091c2b0500034a74d9b65c8717f31bf94dfca552b3ffc7f6fff63aea4037fff67ee4f8d17f7af13ffd9f75fdc5fc5bd03ebdfc21f97f68775e98d1d330b000095e21085408aa263fa21fe0f1f5b02644d14a21326d6b1c9d087293486067e55804f53b0258584433d972144393e713c25b2439442b4fbb219364cb5a11063f00479b9318367100035ef25f3f9318676010317b94999d59a880958f11298c9b59958241ee95444c69b3cf95184201e773936a7f3c412c960a4cd0c9c1273d7d931c37ff6d9d929418b8a94ba43b78934116c9e0bb4c426d6cef2ecc4128781f52a8239071492ce457325d8846086cf25da0f16d8b90c7802ea60d0d0b72d9e6fa8c9f17f2445c39d8fa80800c48e4958dad48ac846d8ad1463e5d72326c95518ff824593090c24165a9e3123b092535d0eb26ca00b30d320b26c8c15b2c03ab49513cb42200013203e03ed995c9b6f64407e93590fe3b91b46fd7e836649f75ee1adc19fb424b45d67997f6a4822f929ef35faddb5df5c53581758c62118564b229411952c81164aae379c084435068ec3223824d712cf19144e250d241082500b5a0c91e191b83c7916a0904a8472259a46942941249d81b785ff89201ea30130a2a474ce62261657760a99323e547108ade54a8e57faefd9549589f0217754f7b64165c0e5ddbc30a69d1b60a8358850e615eaffaf67cfb739b008aebe90e5cdb6599d9411eb29d4bf1193626020bbd249ada1104e39499436342c7e7d9ffd9bb473ee308539a47090c84f69025f2436938b648e12e11aa12260e053c981595e0918abb365111beb127dba50058d03203fe97ea545a7be2a222a70eecada1cf44545fc92fa001c46e917f9f3e4cc0b9b3cf647886082ec8b97fdc800446440b4d7ee15355c86a20181688e56792390929dffe60514948464f344321cdb324ecaacc19dc776c0e59f3bb4d7831c9ae5fed2662cc82f98bdf7ff8659e053949f8f4f8f04bb7bf96d2dcdb3f0af748fb7edfd5b0e0",
		"80e104578424ba31b06db68900101710217c8fcffbffffffff1979833210cc1418840000000125d494cbd56b7a9f71939a407bb23c36490157b3716461bc99a2d721bb5f9322e942220d168209551277cedec9ac2411dde7ddd3f1098064caaae00452acec522331102b6bfa04f7f0f26031a48a98c3bb369e8df22dffd4fd97ffeec0fbb3c6e2b715faf8ef3c6542f86755595f618c773f9e743e6187bab6ae9ce8c9334bdfcef0fc4a22fad25b7d627d0658d1b0787baf71df6df58cd130f1ee8691a96c97c6d78c7696c2b6344b4c5a316341cbd9301eff71baad0e92a58c7f4ebad240852f1a802a7231f4f674e925b06f8efda7744cd6626a102f7e444844b6b10530e512da58e67a4b5823da6f230639396fb448c884a800b66ed7bd39648dceaa5ebc96f5db55e15e311a928a4c14d808318140ac573482c6d952f86ba62071addd4aba2ea4ec411290c5895734ac6aa71a227024b904ba9500f2625574a6ae33a4a3711301219b387e7a1cf996d1d5d21fcfd26d4684d449dd6a1f6edc67af2e0be6e611af2a0c0ad0450128040c82006f93201d01e8b1da15175910c8a08f31785396e7d69fd0115e62d246cc9c64e91c2c9a20232e347629151e210000000049752532f55adea7dc5a0194f8491726a28be359087d659d41678fa573b26eb25664ca0f575139ebf21f8d1bf4f5b5fdee41544239171dff3f4c688fa761594a14ddc40aa522794342bb1dd508bf1b8bb3f8e4a693ecb1cff09c786e084aabc551be968b928cd1519b62b14df3854c6026a84d74a584d3f4ccd9ae7330b4ab0dd1da78a3d355f9ec9718f772945b02f25c34e2db785eb9285f31d74e74003b98948261aaea04c173aae18453f37483fefb0792099cebbd230ee44e059980d83d983025601a7e9aa1a08850830de3e2991b508b305a280e99e6f3b3b6ae65b7e20cd36ba24180e55600cb2b0f56197f41b86417aca1257af9603ec6dd2e569d1c455b9f83c389629e4135e660ee55004e738babd3fedff03562b0896f85c87b7c9ddf0f48b9db3c60c8afcb500380",
		"80e104588424be31b06db689001016d8211c8d2ff7fffffffd1975a43251621000000006af6ba26aa2e3f904d00271b1e487d009e4320471d8f204892ad2252d72b8700260e8b4113e20831048a5a0438f1b93832a048415ff73b5bf69631c8a4e4ca6eb89fc8454ab3558b7ed96e5557083eb2ac47779eb23fa2d680f806dfa4ffd6590df3f6e990132030626063e01e73386af34ea3ff2746f376e6c8fcc3f0574828a15f1375e386521f5b50f0f4aea7ffdb9e0d27caff771bbfd4788dcbd118e539df41f01be771cf33af56f9a35090fd7b57dfd8f557cb9e398b7df5f1ac4d4bb56f69cdd2211b6f49679293c43dc7b283840528ea9f573ee24cc7c055d4b9844491583e315d70af6dd5e73fdab59fd17d532f48dda3cb452cba838dccbbdc2b7d5b4699a8e565a4b71035f2a9bd97bb88b90057c240692aa50aca5486142176a21104b5b19c1a691b086d4118810597934c2302e49091073911ba492badc8dba7ec264c5d29602f779ff3b58a48c510a681f42e7cdfd3ecf37fa7e8f6f8eda069baed02641111cea27687eba9db1f5471d5bb4734a6ca7f63caa22f5c04b5df2a7a5f75000ca648557f5c4ee3bca7e079e382643a1a808cb851d8e67222204200000000d5ed744d545c7f916413a981219ae144e8cafae1383088cb0139c294c240e3f9c9dc1d324047aaf90dd06794b20bebd8f4d73ffefb80f3d6dfde7dddd613573e886de98f79a921d12eb49909e8b4415f66cdc863a11bf1aaada975bd2d91e20d7f58ca2a3219d9bb55d873c2c399a17028a2bdc09a513165845609d47ee11f4f5e8822407a13c1f876c728291c28931eb6004cc5e86a441a29503a82c2b848200ee55e3df0880cc00dbd92d204170376d1d2c19355727b70dd508350425a031cb73d7fee9bbd887878b6b4a3ebb056e59b44b1e7a4884f1e3cfdab2b3607e594e2cf47eb553f54696c0acb08ba8b411ed053572db39a5c885723bc0bd2d97d92b27b5381d5bd555ae9fba09cde8ea4e66a7bc1b3488deebff89ad1d8aac49546b79e301c0",
	}

	testHelperTemplete(t, base.AvPacketPtAac, 32000, 128, ss, func(rtpPackets []RtpPacket) []base.AvPacket {
		return []base.AvPacket{
			{
				Timestamp:   69281105,
				PayloadType: base.AvPacketPtAac,
				Payload:     rtpPackets[0].Raw[12+2+6 : 12+2+6+24],
			},
			{
				Timestamp:   69281137,
				PayloadType: base.AvPacketPtAac,
				Payload:     rtpPackets[0].Raw[12+2+6+24 : 12+2+6+24+6],
			},
			{
				Timestamp:   69281169,
				PayloadType: base.AvPacketPtAac,
				Payload:     rtpPackets[0].Raw[12+2+6+24+6:],
			},
			{
				Timestamp:   69281201,
				PayloadType: base.AvPacketPtAac,
				Payload:     rtpPackets[1].Raw[12+4:],
			},
			{
				Timestamp:   69281233,
				PayloadType: base.AvPacketPtAac,
				Payload:     rtpPackets[2].Raw[12+4:],
			},
			{
				Timestamp:   69281265,
				PayloadType: base.AvPacketPtAac,
				Payload:     rtpPackets[3].Raw[12+4:],
			},
		}
	})
}

// testHelperTemplete
//
// @param hexRtpPackets: rtp包的二进制数组
// @param expectedFn:
// []RtpPacket: `hexRtpPackets`解析成的rtp包数组
// []base.AvPacket: rtp包数组解析成的AvPacket数组
func testHelperTemplete(t *testing.T, payloadType base.AvPacketPt, clockRate int, maxSize int, hexRtpPackets []string, expectedFn func([]RtpPacket) []base.AvPacket) {
	rtpPackets, err := testHelperHexstream2rtppackets(hexRtpPackets)
	assert.Equal(t, nil, err)

	expected := expectedFn(rtpPackets)
	assert.Equal(t, expected, testHelperUnpack(payloadType, clockRate, maxSize, rtpPackets))
}

func testHelperAddPrefixLength(in []byte) (out []byte) {
	out = make([]byte, len(in)+4)
	bele.BePutUint32(out, uint32(len(in)))
	copy(out[4:], in)
	return
}

// ---------------------------------------------------------------------------------------------------------------------

func testHelperUnpack(payloadType base.AvPacketPt, clockRate int, maxSize int, rtpPackets []RtpPacket) []base.AvPacket {
	var outPkts []base.AvPacket
	unpacker := DefaultRtpUnpackerFactory(payloadType, clockRate, maxSize, func(pkt base.AvPacket) {
		Log.Debugf("%s", hex.EncodeToString(pkt.Payload))
		outPkts = append(outPkts, pkt)
	})
	for _, pkt := range rtpPackets {
		unpacker.Feed(pkt)
	}
	return outPkts
}

func testHelperHexstream2rtppackets(hexPackets []string) (pkts []RtpPacket, err error) {
	var pkt RtpPacket
	for _, p := range hexPackets {
		pkt, err = testHelperHexstream2rtppacket(p)
		if err != nil {
			return
		}
		pkts = append(pkts, pkt)
	}
	return
}

func testHelperHexstream2rtppacket(hexPacket string) (pkt RtpPacket, err error) {
	var raw []byte
	raw, err = hex.DecodeString(hexPacket)
	if err != nil {
		return
	}
	pkt, err = ParseRtpPacket(raw)
	return
}
