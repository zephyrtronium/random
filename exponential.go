package random

import "math"

func genExpo(rng RNG) float64 {
	for {
		x := rng.Uint64()
		j := uint32(x)
		i := int(x >> 32 & 255)
		if j < expoK[i] {
			return float64(j) * float64(expoW[i])
		}
		if i != 0 {
			v := float64(j) * float64(expoW[i])
			if expoF[i]+rng.Float32()*(expoF[i-1]-expoF[i]) < float32(math.Exp(-v)) {
				return v
			}
		} else {
			return expoR - math.Log(rng.Float64())
		}
	}
}

const expoR = 7.697117470131049714044628048

var expoK = [256]uint32{
	0xe290a139, 0x0, 0x9beadebc, 0xc377ac71, 0xd4ddb990, 0xde893fb8,
	0xe4a8e87c, 0xe8dff16a, 0xebf2deab, 0xee49a6e8, 0xf0204efd, 0xf19bdb8e,
	0xf2d458bb, 0xf3da104b, 0xf4b86d78, 0xf577ad8a, 0xf61de83d, 0xf6afb784,
	0xf730a573, 0xf7a37651, 0xf80a5bb6, 0xf867189d, 0xf8bb1b4f, 0xf9079062,
	0xf94d70ca, 0xf98d8c7d, 0xf9c8928a, 0xf9ff175b, 0xfa319996, 0xfa6085f8,
	0xfa8c3a62, 0xfab5084e, 0xfadb36c8, 0xfaff0410, 0xfb20a6ea, 0xfb404fb4,
	0xfb5e2951, 0xfb7a59e9, 0xfb95038c, 0xfbae44ba, 0xfbc638d8, 0xfbdcf892,
	0xfbf29a30, 0xfc0731df, 0xfc1ad1ed, 0xfc2d8b02, 0xfc3f6c4d, 0xfc5083ac,
	0xfc60ddd1, 0xfc708662, 0xfc7f8810, 0xfc8decb4, 0xfc9bbd62, 0xfca9027c,
	0xfcb5c3c3, 0xfcc20864, 0xfccdd70a, 0xfcd935e3, 0xfce42ab0, 0xfceebace,
	0xfcf8eb3b, 0xfd02c0a0, 0xfd0c3f59, 0xfd156b7b, 0xfd1e48d6, 0xfd26daff,
	0xfd2f2552, 0xfd372af7, 0xfd3eeee5, 0xfd4673e7, 0xfd4dbc9e, 0xfd54cb85,
	0xfd5ba2f2, 0xfd62451b, 0xfd68b415, 0xfd6ef1da, 0xfd750047, 0xfd7ae120,
	0xfd809612, 0xfd8620b4, 0xfd8b8285, 0xfd90bcf5, 0xfd95d15e, 0xfd9ac10b,
	0xfd9f8d36, 0xfda43708, 0xfda8bf9e, 0xfdad2806, 0xfdb17141, 0xfdb59c46,
	0xfdb9a9fd, 0xfdbd9b46, 0xfdc170f6, 0xfdc52bd8, 0xfdc8ccac, 0xfdcc542d,
	0xfdcfc30b, 0xfdd319ef, 0xfdd6597a, 0xfdd98245, 0xfddc94e5, 0xfddf91e6,
	0xfde279ce, 0xfde54d1f, 0xfde80c52, 0xfdeab7de, 0xfded5034, 0xfdefd5be,
	0xfdf248e3, 0xfdf4aa06, 0xfdf6f984, 0xfdf937b6, 0xfdfb64f4, 0xfdfd818d,
	0xfdff8dd0, 0xfe018a08, 0xfe03767a, 0xfe05536c, 0xfe07211c, 0xfe08dfc9,
	0xfe0a8fab, 0xfe0c30fb, 0xfe0dc3ec, 0xfe0f48b1, 0xfe10bf76, 0xfe122869,
	0xfe1383b4, 0xfe14d17c, 0xfe1611e7, 0xfe174516, 0xfe186b2a, 0xfe19843e,
	0xfe1a9070, 0xfe1b8fd6, 0xfe1c8289, 0xfe1d689b, 0xfe1e4220, 0xfe1f0f26,
	0xfe1fcfbc, 0xfe2083ed, 0xfe212bc3, 0xfe21c745, 0xfe225678, 0xfe22d95f,
	0xfe234ffb, 0xfe23ba4a, 0xfe241849, 0xfe2469f2, 0xfe24af3c, 0xfe24e81e,
	0xfe25148b, 0xfe253474, 0xfe2547c7, 0xfe254e70, 0xfe25485a, 0xfe25356a,
	0xfe251586, 0xfe24e88f, 0xfe24ae64, 0xfe2466e1, 0xfe2411df, 0xfe23af34,
	0xfe233eb4, 0xfe22c02c, 0xfe22336b, 0xfe219838, 0xfe20ee58, 0xfe20358c,
	0xfe1f6d92, 0xfe1e9621, 0xfe1daef0, 0xfe1cb7ac, 0xfe1bb002, 0xfe1a9798,
	0xfe196e0d, 0xfe1832fd, 0xfe16e5fe, 0xfe15869d, 0xfe141464, 0xfe128ed3,
	0xfe10f565, 0xfe0f478c, 0xfe0d84b1, 0xfe0bac36, 0xfe09bd73, 0xfe07b7b5,
	0xfe059a40, 0xfe03644c, 0xfe011504, 0xfdfeab88, 0xfdfc26e9, 0xfdf98629,
	0xfdf6c83b, 0xfdf3ec01, 0xfdf0f04a, 0xfdedd3d1, 0xfdea953d, 0xfde7331e,
	0xfde3abe9, 0xfddffdfb, 0xfddc2791, 0xfdd826cd, 0xfdd3f9a8, 0xfdcf9dfc,
	0xfdcb1176, 0xfdc65198, 0xfdc15bb3, 0xfdbc2ce2, 0xfdb6c206, 0xfdb117be,
	0xfdab2a63, 0xfda4f5fd, 0xfd9e7640, 0xfd97a67a, 0xfd908192, 0xfd8901f2,
	0xfd812182, 0xfd78d98e, 0xfd7022bb, 0xfd66f4ed, 0xfd5d4732, 0xfd530f9c,
	0xfd48432b, 0xfd3cd59a, 0xfd30b936, 0xfd23dea4, 0xfd16349e, 0xfd07a7a3,
	0xfcf8219b, 0xfce7895b, 0xfcd5c220, 0xfcc2aadb, 0xfcae1d5e, 0xfc97ed4e,
	0xfc7fe6d4, 0xfc65ccf3, 0xfc495762, 0xfc2a2fc8, 0xfc07ee19, 0xfbe213c1,
	0xfbb8051a, 0xfb890078, 0xfb5411a5, 0xfb180005, 0xfad33482, 0xfa839276,
	0xfa263b32, 0xf9b72d1c, 0xf930a1a2, 0xf889f023, 0xf7b577d2, 0xf69c650c,
	0xf51530f0, 0xf2cb0e3c, 0xeeefb15d, 0xe6da6ecf,
}

var expoW = [256]float32{
	2.02495545850393e-9, 1.48667403997396e-11, 2.4409617196261e-11,
	3.19688070891457e-11, 3.84467706466532e-11, 4.42282039724367e-11,
	4.95164447070489e-11, 5.44335886509334e-11, 5.90594400153292e-11,
	6.34494203791175e-11, 6.76438108764661e-11, 7.16729449748371e-11,
	7.55603231994692e-11, 7.93245809769374e-11, 8.29807855790468e-11,
	8.65413214382524e-11, 9.00165126521886e-11, 9.34150719308011e-11,
	9.67444315553543e-11, 1.00010992080302e-10, 1.03220312407602e-10,
	1.06377257251046e-10, 1.09486113088711e-10, 1.12550680444916e-10,
	1.15574348140199e-10, 1.18560153628619e-10, 1.2151083247553e-10,
	1.24428859268587e-10, 1.27316481704663e-10, 1.30175749191908e-10,
	1.33008537006702e-10, 1.35816566820436e-10, 1.38601424240392e-10,
	1.41364573878306e-10, 1.44107372359111e-10, 1.4683107960352e-10,
	1.49536868656179e-10, 1.52225834282037e-10, 1.54899000514457e-10,
	1.57557327307184e-10, 1.60201716416923e-10, 1.62833016622633e-10,
	1.65452028370848e-10, 1.68059507922446e-10, 1.70656171064909e-10,
	1.73242696444622e-10, 1.75819728565864e-10, 1.78387880496549e-10,
	1.80947736315227e-10, 1.83499853329149e-10, 1.86044764089279e-10,
	1.88582978224712e-10, 1.91114984116147e-10, 1.93641250425548e-10,
	1.96162227497056e-10, 1.98678348642395e-10, 2.01190031322419e-10,
	2.03697678235133e-10, 2.06201678319311e-10, 2.08702407681823e-10,
	2.11200230455885e-10, 2.13695499596662e-10, 2.16188557619977e-10,
	2.18679737289265e-10, 2.2116936225539e-10, 2.23657747653468e-10,
	2.2614520066043e-10, 2.28632021016689e-10, 2.31118501514959e-10,
	2.3360492845897e-10, 2.36091582094575e-10, 2.38578737015514e-10,
	2.41066662545905e-10, 2.43555623101314e-10, 2.46045878530143e-10,
	2.4853768443688e-10, 2.51031292488652e-10, 2.5352695070639e-10,
	2.56024903741804e-10, 2.58525393141297e-10, 2.61028657597799e-10,
	2.6353493319151e-10, 2.66044453620369e-10, 2.68557450421102e-10,
	2.71074153181556e-10, 2.73594789745033e-10, 2.76119586407254e-10,
	2.78648768106569e-10, 2.81182558607953e-10, 2.83721180681323e-10,
	2.8626485627467e-10, 2.88813806682454e-10, 2.91368252709706e-10,
	2.93928414832245e-10, 2.96494513353389e-10, 2.99066768557535e-10,
	3.01645400860952e-10, 3.04230630960123e-10, 3.0682267997794e-10,
	3.09421769608072e-10, 3.12028122257792e-10, 3.14641961189531e-10,
	3.17263510661453e-10, 3.19892996067295e-10, 3.22530644075741e-10,
	3.25176682769563e-10, 3.27831341784805e-10, 3.30494852450207e-10,
	3.33167447927147e-10, 3.35849363350312e-10, 3.38540835969335e-10,
	3.41242105291631e-10, 3.43953413226673e-10, 3.46675004231917e-10,
	3.4940712546064e-10, 3.52150026911897e-10, 3.5490396158286e-10,
	3.57669185623767e-10, 3.60445958495725e-10, 3.63234543131638e-10,
	3.66035206100491e-10, 3.68848217775274e-10, 3.71673852504809e-10,
	3.7451238878976e-10, 3.77364109463118e-10, 3.80229301875455e-10,
	3.83108258085261e-10, 3.86001275054685e-10, 3.88908654851013e-10,
	3.91830704854232e-10, 3.94767737971045e-10, 3.97720072855721e-10,
	4.00688034138161e-10, 4.0367195265963e-10, 4.0667216571655e-10,
	4.09689017312851e-10, 4.12722858421343e-10, 4.15774047254614e-10,
	4.1884294954601e-10, 4.21929938841236e-10, 4.25035396801196e-10,
	4.28159713516682e-10, 4.313032878356e-10, 4.34466527703411e-10,
	4.3764985051756e-10, 4.40853683496664e-10, 4.44078464065303e-10,
	4.47324640255311e-10, 4.50592671124509e-10, 4.53883027193878e-10,
	4.57196190904255e-10, 4.60532657093685e-10, 4.63892933496641e-10,
	4.67277541266409e-10, 4.70687015522021e-10, 4.74121905921206e-10,
	4.77582777260939e-10, 4.81070210107271e-10, 4.84584801456245e-10,
	4.88127165427831e-10, 4.91697933994942e-10, 4.95297757749764e-10,
	4.98927306709774e-10, 5.02587271166007e-10, 5.06278362576332e-10,
	5.10001314506684e-10, 5.13756883623466e-10, 5.17545850740521e-10,
	5.21369021924424e-10, 5.25227229662058e-10, 5.29121334094823e-10,
	5.33052224324148e-10, 5.37020819793357e-10, 5.41028071751398e-10,
	5.4507496480435e-10, 5.49162518561198e-10, 5.53291789380866e-10,
	5.57463872228157e-10, 5.61679902646893e-10, 5.65941058859327e-10,
	5.70248564001697e-10, 5.74603688506727e-10, 5.79007752644878e-10,
	5.83462129237269e-10, 5.8796824655445e-10, 5.92527591416582e-10,
	5.971417125121e-10, 6.01812223953693e-10, 6.06540809092306e-10,
	6.11329224612049e-10, 6.16179304931268e-10, 6.21092966937755e-10,
	6.26072215089063e-10, 6.31119146912342e-10, 6.3623595894191e-10,
	6.41424953137139e-10, 6.46688543828148e-10, 6.52029265242335e-10,
	6.57449779671159e-10, 6.62952886343745e-10, 6.68541531082135e-10,
	6.74218816822428e-10, 6.79988015096807e-10, 6.85852578583883e-10,
	6.91816154849038e-10, 6.97882601412975e-10, 7.04056002305745e-10,
	7.10340686285741e-10, 7.16741246928948e-10, 7.23262564823922e-10,
	7.29909832143327e-10, 7.36688579904375e-10, 7.43604708279539e-10,
	7.50664520376889e-10, 7.57874759978254e-10, 7.65242653805546e-10,
	7.72775958983868e-10, 7.80483016488168e-10, 7.88372811502847e-10,
	7.96455041796696e-10, 8.04740195426336e-10, 8.13239639339517e-10,
	8.21965720767468e-10, 8.30931883689095e-10, 8.40152803139973e-10,
	8.49644540753414e-10, 8.59424725695844e-10, 8.6951276614326e-10,
	8.79930097705607e-10, 8.90700476831369e-10, 9.0185032933939e-10,
	9.13409167000905e-10, 9.25410088774233e-10, 9.37890388222396e-10,
	9.50892295317794e-10, 9.64463889986289e-10, 9.786602374481e-10,
	9.93544813310114e-10, 1.00919131196972e-9, 1.02568596915192e-9,
	1.04313058464984e-9, 1.06164651496973e-9, 1.08138003512753e-9,
	1.10250967475626e-9, 1.12525647064324e-9, 1.14989864777337e-9,
	1.17679324233469e-9, 1.20640901878977e-9, 1.2393785886826e-9,
	1.27658495389066e-9, 1.31931392649515e-9, 1.36954344711159e-9,
	1.43054981384717e-9, 1.50836503455242e-9, 1.61608532755105e-9,
	1.79212481485006e-9,
}

var expoF = [256]float32{
	1.0, 0.938143680862175, 0.900469929925746,
	0.871704332381204, 0.84778550062399, 0.82699329664305,
	0.808421651523008, 0.791527636972496, 0.775956852040116,
	0.761463388849896, 0.747868621985195, 0.735038092431424,
	0.722867659593572, 0.711274760805076, 0.700192655082788,
	0.689566496117078, 0.679350572264765, 0.669506316731925,
	0.660000841079, 0.650805833414571, 0.641896716427266,
	0.633251994214366, 0.624852738703666, 0.616682180915208,
	0.608725382079622, 0.600968966365232, 0.593400901691733,
	0.586010318477268, 0.578787358602845, 0.571723048664826,
	0.5648091929124, 0.558038282262587, 0.551403416540641,
	0.54489823767244, 0.538516872002862, 0.532253880263043,
	0.52610421398362, 0.520063177368234, 0.514126393814749,
	0.508289776410643, 0.502549501841348, 0.49690198724155,
	0.491343869594033, 0.485871987341885, 0.480483363930454,
	0.475175193037377, 0.46994482528396, 0.464789756250426,
	0.459707615642138, 0.454696157474615, 0.449753251162755,
	0.444876873414549, 0.440065100842354, 0.435316103215637,
	0.430628137288459, 0.425999541143034, 0.421428728997617,
	0.416914186433003, 0.412454465997161, 0.408048183152032,
	0.40369401253053, 0.399390684475231, 0.39513698183329,
	0.390931736984797, 0.386773829084138, 0.38266218149601,
	0.378595759409581, 0.374573567615902, 0.370594648435146,
	0.366658079781514, 0.362762973354818, 0.35890847294875,
	0.355093752866787, 0.351318016437483, 0.347580494621637,
	0.343880444704502, 0.34021714906678, 0.336589914028678,
	0.332998068761809, 0.329440964264136, 0.325917972393556,
	0.322428484956089, 0.318971912844957, 0.315547685227129,
	0.31215524877418, 0.30879406693456, 0.30546361924459,
	0.302163400675694, 0.298892921015582, 0.295651704281261,
	0.292439288161893, 0.289255223489678, 0.286099073737077,
	0.282970414538781, 0.279868833236973, 0.276793928448517,
	0.273745309652803, 0.27072259679906, 0.267725419932045,
	0.264753418835062, 0.261806242689363, 0.258883549749016,
	0.255985007030415, 0.253110290015629, 0.250259082368862,
	0.247431075665328, 0.244625969131892, 0.241843469398877,
	0.239083290262449, 0.23634515245706, 0.233628783437433,
	0.230933917169627, 0.228260293930717, 0.225607660116684,
	0.22297576805812, 0.220364375843359, 0.217773247148701,
	0.215202151075379, 0.212650861992978, 0.210119159388988,
	0.207606827724222, 0.205113656293838, 0.202639439093709,
	0.200183974691911, 0.197747066105099, 0.195328520679563,
	0.192928149976771, 0.190545769663195, 0.188181199404254,
	0.185834262762197, 0.183504787097767, 0.181192603475496,
	0.178897546572478, 0.176619454590495, 0.174358169171353,
	0.17211353531532, 0.169885401302528, 0.16767361861725,
	0.165478041874936, 0.163298528751902, 0.161134939917592,
	0.158987138969314, 0.156854992369365, 0.154738369384468,
	0.152637142027443, 0.15055118500104, 0.148480375643867,
	0.146424593878345, 0.144383722160635, 0.142357645432472,
	0.140346251074862, 0.13834942886358, 0.136367070926429,
	0.134399071702214, 0.132445327901388, 0.130505738468331,
	0.128580204545228, 0.126668629437511, 0.124770918580831,
	0.122886979509545, 0.121016721826675, 0.119160057175328,
	0.117316899211556, 0.115487163578634, 0.113670767882744,
	0.111867631670056, 0.110077676405185, 0.108300825451034,
	0.106537004050002, 0.10478613930657, 0.103048160171258,
	0.101322997425954, 0.0996105836706371, 0.0979108533114922,
	0.0962237425504328, 0.0945491893760559, 0.0928871335560435,
	0.0912375166310402, 0.0896002819100329, 0.0879753744672702,
	0.0863627411407569, 0.0847623305323681, 0.0831740930096324,
	0.0815979807092374, 0.0800339475423199, 0.0784819492016064,
	0.0769419431704805, 0.0754138887340584, 0.0738977469923647,
	0.0723934808757087, 0.0709010551623718, 0.0694204364987288,
	0.0679515934219366, 0.0664944963853398, 0.0650491177867538,
	0.0636154319998073, 0.062193415408541, 0.0607830464454796,
	0.0593843056334203, 0.0579971756312007, 0.0566216412837429,
	0.055257689676697, 0.0539053101960461, 0.0525644945930717,
	0.0512352370551263, 0.0499175342827064, 0.0486113855733795,
	0.0473167929131815, 0.0460337610761752, 0.0447622977329433,
	0.0435024135688882, 0.0422541224133162, 0.0410174413804148,
	0.0397923910233741, 0.0385789955030749, 0.0373772827729594,
	0.0361872847819314, 0.0350090376973974, 0.0338425821508743,
	0.0326879635089595, 0.0315452321728936, 0.0304144439104666,
	0.0292956602246374, 0.0281889487639786, 0.0270943837809558,
	0.0260120466451342, 0.0249420264197318, 0.0238844205115582,
	0.0228393354063852, 0.0218068875042836, 0.0207872040725781,
	0.0197804243380097, 0.018786700744696, 0.0178062004109114,
	0.0168391068260399, 0.0158856218399732, 0.0149459680116911,
	0.0140203914031819, 0.013109164931255, 0.0122125924262554,
	0.0113310135978346, 0.01046481018103, 0.00961441364250221,
	0.00878031498580898, 0.00796307743801704, 0.00716335318363498,
	0.00638190593731918, 0.00561964220720548, 0.00487765598354239,
	0.0041572951208338, 0.0034602647778369, 0.00278879879357408,
	0.00214596774371891, 0.00153629978030157, 0.000967269282327175,
	0.000454134353841497,
}