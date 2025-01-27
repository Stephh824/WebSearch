package main

import (
	"testing"
	"fmt"
)

type TfIdfCase struct {
	name string
	term string
	want Hits
}

type Hit struct {
	url string
	tfidf float64
}

type Hits []Hit

func TestCase1(t *testing.T) {
	tcs := []TfIdfCase {
		{
			name: "bigram test",
			term: "artificial intelligence",
			want: Hits {
				{"Will Hurd joins OpenAI’s board of directors", 0.06027525273978421},
				{"Research Scientist", 0.030392597828182225},
				{"Media Relations, Europe Lead", 0.02952338151948668},
				{"Organizational update from OpenAI", 0.026414743112434847},
				{"Senior Manager, Strategic Sourcing (Non-Technology)", 0.025999623080064076},
				{"Software Engineer, Safety", 0.025263045452117713},
				{"Media Relations, Corporate Communications", 0.024093930672643458},
				{"Research Engineer - Data Specialization", 0.023161863722057634},
				{"Research Engineer, Privacy", 0.022508803654706384},
				{"Research Engineer, Post Training Infra", 0.022175339896858885},
				{"Research Engineer - Fine-Tuning API", 0.02201228592702904},
				{"Software Engineer – Model Inference", 0.021332571634745483},
				{"OpenAI customer story: Yabble", 0.02113179448994788},
				{"Introducing OpenAI", 0.021082189338563023},
				{"About", 0.021082189338563023},
				{"US Congressional Lead, Global Affairs", 0.021045138039198236},
				{"Software Engineer, Privacy", 0.020646006110868614},
				{"Confidence-Building Measures for Artificial Intelligence: Workshop proceedings", 0.02059865288584369},
				{"International Payroll Specialist", 0.02041139240506329},
				{"Security Engineer, Detection and Response", 0.02027316627139469},
				{"Strategic Finance - Infrastructure", 0.02026173188545482},
				{"Security Engineer, Detection & Response", 0.019957805907173},
				{"Senior Software Engineer- Identity Platform", 0.019913553565915405},
				{"Software Engineer, Model Inference", 0.019781966207550326},
				{"Engineering Manager – Fine Tuning API", 0.019738489358742523},
				{"Software Engineer, Privacy", 0.019652106473146275},
				{"Software Engineer, Security Data Platform", 0.019609197943728923},
				{"Senior Revenue Accountant, Deal Desk", 0.019159493670886076},
				{"Partnership with American Journalism Project to support local news", 0.019098378858538753},
				{"Manager, Enterprise Platform Sales", 0.01891735157078009},
				{"Media Relations, Europe Lead", 0.018867673651739178},
				{"Account Associate", 0.018366079055680672},
				{"Assistant Controller", 0.01817099172125007},
				{"Security Engineer, Partnerships", 0.01785489594081083},
				{"Model Teacher (Contract)", 0.017714028911692006},
				{"DALL·E 3 system card", 0.017106690777576854},
				{"Research Engineer", 0.01686575147085042},
				{"Commitment to diversity, equity & inclusion", 0.016182004789599726},
				{"OpenAI’s technology explained", 0.015909676985346056},
				{"Workplace Triage Coordinator", 0.01571480780092362},
				{"Systems Software Engineer, Frontiers", 0.015632746141388768},
				{"OpenAI announces leadership transition", 0.015592035864978902},
				{"Software Engineer, Developer Experience", 0.015391624092935471},
				{"Global Compensation Manager", 0.015170629490249743},
				{"Workplace Specialist", 0.015119549929676512},
				{"Report from the OpenAI hackathon", 0.015081465421037528},
				{"Software Engineer, Distributed Systems", 0.014968354430379747},
				{"Software Engineer, Leverage Engineering", 0.014856927474322329},
				{"Software Engineer, Developer Productivity", 0.014674857284686026},
				{"Technical Program Manager, Applied Infrastructure", 0.014603272615004631},
				{"IT Support", 0.014567741538082478},
				{"ML User Experience Engineer", 0.014532382942116258},
				{"GPT-4V(ision) system card", 0.014532382942116258},
				{"Software Engineer, Fullstack", 0.01442732957145036},
				{"DL HW/SW Codesign Engineer", 0.01442732957145036},
				{"Graph Compiler Engineer", 0.014323784143904063},
				{"Software Engineer, Triton Compiler", 0.014255575647980712},
				{"Research Software Engineer, Data Quality", 0.014255575647980712},
				{"Engineering Manager, DALL-E", 0.014255575647980712},
				{"Software Engineer, Backend", 0.014154472274590777},
				{"Research Program Manager, Basic Research", 0.014154472274590777},
				{"Distributed Systems/ML Engineer", 0.014154472274590777},
				{"Software Engineer, Full Stack (DALL-E)", 0.014087862993298584},
				{"AI Policy Counsel", 0.014054792892375349},
				{"Software Engineer, Infrastructure", 0.014021877686538405},
				{"Senior Software Engineer, Data Acquisition", 0.013924050632911392},
				{"Applied AI: Senior Software Engineer, iOS Growth", 0.01376400407391241},
				{"Understanding the capabilities, limitations, and societal impact of large language models", 0.013732435257229126},
				{"Applied AI: Senior Software Engineer, Android Growth", 0.013638591736109108},
				{"Full-Stack Engineer, Communications & Design", 0.013424533121416813},
				{"Data Scientist, Product", 0.013334836909024273},
				{"Associate General Counsel, Regulatory", 0.01324633135431836},
				{"Research Scientist, Safety", 0.01310140431543085},
				{"Data Infrastructure Engineer", 0.013015960374243258},
				{"Software Engineer, Backend (New Products)", 0.012987726186880474},
				{"Head of Commercial Legal", 0.012876003811079352},
				{"Software Engineer, Anti Fraud & Abuse", 0.012766187147445414},
				{"Research Engineer, Preparedness", 0.012766187147445414},
				{"Account Engineer", 0.012766187147445414},
				{"User Operations Generalist", 0.012739025047131699},
				{"Kernel Engineer", 0.012658227848101266},
				{"Senior Manager of Procure to Pay Operations", 0.012631522726058856},
				{"Our structure", 0.01257844910115945},
				{"Senior AI Product Counsel", 0.012552079186901256},
				{"Summarizing books with human feedback", 0.01244769599199979},
				{"Developer Engagement Specialist (Contract)", 0.012421870896580702},
				{"National Security Threat Researcher", 0.012396152737374532},
				{"Data Engineer, Applied Engineering", 0.012345034581756493},
				{"Research Scientist, Machine Learning", 0.012269142975721105},
				{"Engineering Manager, AI Inference Systems", 0.012194178762020159},
				{"OpenAI customer story: Duolingo", 0.012169393845837194},
				{"OpenAI supporters", 0.0121201250448419},
				{"OpenAI", 0.0121201250448419},
				{"OpenAI", 0.0121201250448419},
				{"OpenAI", 0.0121201250448419},
				{"OpenAI", 0.0121201250448419},
				{"Solutions Architect", 0.012046965336321729},
				{"Solutions Architect, London", 0.012022774642875298},
				{"Software Engineer, Infrastructure", 0.012022774642875298},
				{"FullStack Software Engineer, Growth", 0.011879646373317258},
				{"Research Engineer, Superalignment", 0.01167123152466257},
				{"Software Engineer — Engineering Acceleration", 0.01164852484854455},
				{"Backend Software Engineer, Growth", 0.011625906353693007},
				{"Sales Engineer, Dublin", 0.011603375527426161},
				{"Research Scientist, Superalignment", 0.011603375527426161},
				{"Workplace Coordinator", 0.011536303992585546},
				{"Stream Infrastructure Engineer", 0.011470003394927009},
				{"HW/SW Co-design Engineer", 0.011318226412385442},
				{"IT Network Engineer", 0.011212250509647751},
				{"Machine Learning Engineer, Moderation", 0.011087669948429442},
				{"Sales Engineer", 0.01100614296351452},
				{"OpenAI leadership team update", 0.010965827421523624},
				{"OpenAI Fellows Fall 2018", 0.010578342353625262},
				{"OpenAI Fellows Winter 2019 & Interns Summer 2019", 0.010522569019599118},
				{"Introducing OpenAI London", 0.010504108372196314},
				{"Safety & responsibility", 0.0103946905766526},
				{"Reliability Engineer", 0.00967260383223247},
				{"OpenAI Charter", 0.009297114553030898},
				{"Announcing OpenAI’s Bug Bounty Program", 0.009239724957024534},
				{"", 0.009062575840795003},
				{"Introducing OpenAI Dublin", 0.009017080982156474},
				{"A research agenda for assessing the economic impacts of code generation models", 0.008753423643496927},
				{"Microsoft invests in and partners with OpenAI to support us building beneficial AGI", 0.007961890654457312},
				{"OpenAI interview guide", 0.007878081279147235},
				{"Planning for AGI and beyond", 0.007391779965619628},
				{"Team update", 0.007043931496649292},
				{"Research", 0.007043931496649292},
				{"Careers", 0.007043931496649292},
				{"OpenAI Residency", 0.006682301084990958},
				{"AI safety needs social scientists", 0.006637851188638469},
				{"Procgen and MineRL Competitions", 0.006479807112718505},
				{"DALL·E API now available in public beta", 0.006479807112718505},
				{"Economic impacts research at OpenAI", 0.006047819971870605},
				{"Security", 0.005975390990171555},
				{"OpenAI customer story: Stripe", 0.005858455745745498},
				{"EU terms of use", 0.005850822578650064},
				{"Terms of use", 0.005581735648370322},
				{"Learning to cooperate, compete, and communicate", 0.00504834888039789},
				{"OpenAI LP", 0.004782221862741133},
				{"Team update", 0.0045774784190763754},
				{"Frontier Model Forum updates", 0.004481543242628667},
				{"OpenAI Scholars 2018: Meet our Scholars", 0.0043895467537770514},
				{"The power of continuous learning", 0.004240327034101912},
				{"GPT-2: 6-month follow-up", 0.004050975488600743},
				{"OpenAI API", 0.004034596881503975},
				{"How should AI systems behave, and who should decide?", 0.003954651104459642},
				{"Business terms - August 2023", 0.003930421294629255},
				{"Forecasting potential misuses of language models for disinformation campaigns and how to reduce risk", 0.003913295275916273},
				{"Business terms", 0.0037467720726857942},
				{"Our approach to alignment research", 0.0033300009856239705},
				{"OpenAI customer story: Government of Iceland", 0.0031119239979999474},
				{"OpenAI Supplier Code of Conduct", 0.0030516522793842503},
				{"Introducing ChatGPT and Whisper APIs", 0.002719047126317847},
				{"Learning dexterity", 0.0023834959284044182},
				{"Research index", 0.0020876366011687232},
				{"Terms of use - March 2023", 0.0015739594564016558},
				{"Universe", 0.001453238294211626},
				{"Better language models and their implications", 0.0010482040917632876},			
			},
		},
	}
	fmt.Println("hello")
	indx := indexInIt()

	indx.Open()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if urls, found := dlookup(tc.term, &indx, false); found {
				for idx, ti := range urls {
					if ti.url != tc.want[idx].url || ti.tfidf != tc.want[idx].tfidf {
						t.Errorf("%v failed\nGot: %v: %v Expected %v: %v\n", tc.name, ti.url, ti.tfidf, tc.want[idx].url, tc.want[idx].tfidf)
					}
				}
			}
		})
	}
	indx.Close()
}


func TestCas2(t *testing.T) {
	tcs := []TfIdfCase {
		{
			name: "wildcard test",
			term: "rob",
			want: Hits {
				{"Roboschool", 4.936432637571158},
				{"Roboschool", 1.7950664136622392},
				{"Learning with opponent-learning awareness", 1.5508196721311476},
				{"Proximal Policy Optimization", 1.2513227513227512},
				{"OpenAI Robotics Symposium 2019", 0.41932624113475175},
				{"OpenAI Robotics Symposium 2019", 0.40055043929289724},
				{"Computational limitations in robust classification and win-win results", 0.3494385342789598},
				{"Solving math word problems", 0.3384725474889409},
				{"OpenAI Scholars 2019: Meet our Scholars", 0.3248626373626374},
				{"Generalizing from simulation", 0.24250882989631994},
				{"Image GPT", 0.24194373401534525},
				{"Universe", 0.22961165048543689},
				{"DALL·E 3", 0.22653256704980843},
				{"Testing robustness against unforeseen adversaries", 0.2233814185127941},
				{"Transfer of adversarial robustness between perturbation types", 0.2109454638124363},
				{"Jukebox", 0.17957479119210326},
				{"Solving Rubik’s Cube with a robot hand", 0.17218784128139789},
				{"Robots that learn", 0.1673636129696346},
				{"Confidence-Building Measures for Artificial Intelligence: Workshop proceedings", 0.15498034076015726},
				{"Multi-Goal Reinforcement Learning: Challenging robotics environments and request for research", 0.15481801518722177},
				{"Roboschool", 0.14735619813645248},
				{"AI-written critiques help humans notice flaws", 0.14689440993788822},
				{"Asymmetric actor critic for image-based robot learning", 0.1452613475830723},
				{"Sim-to-real transfer of robotic control with dynamics randomization", 0.14233269138180069},
				{"GPT-4V(ision) technical work and authors", 0.1365473441108545},
				{"Frontier AI regulation: Managing emerging risks to public safety", 0.1362327188940092},
				{"Learning dexterity", 0.129277973191368},
				{"Universe", 0.11480582524271844},
				{"Ingredients for robotics research", 0.11289501320688668},
				{"Safety Gym", 0.11269762929934138},
				{"Retro Contest", 0.1042768959435626},
				{"Domain randomization and generative models for robotic grasping", 0.10194514790667601},
				{"Robust adversarial inputs", 0.09746950214309265},
				{"OpenAI Scholars 2019: Meet our Scholars", 0.09281789638932497},
				{"OpenAI technical goals", 0.08715680854984337},
				{"Faster physics in Python", 0.08505664448840138},
				{"Research index", 0.08246164574616457},
				{"Researcher Access Program application", 0.08036017669045192},
				{"Transfer from simulation to real world through learning deep inverse dynamics model", 0.07815167703915844},
				{"OpenAI Gym Beta", 0.07418250254855321},
				{"Spam detection in the physical world", 0.073031394750386},
				{"Welcome, Pieter and Shivon!", 0.06723525230987919},
				{"Learning Day", 0.06434623756218906},
				{"Hindsight Experience Replay", 0.063600914347183},
				{"Introducing Whisper", 0.06082818930041152},
				{"OpenAI Scholars 2019: Final projects", 0.06027781317701032},
				{"Sharing & publication policy", 0.05772798281585628},
				{"Research Scientist, Safety", 0.0575006078288354},
				{"Concrete AI safety problems", 0.054292929292929296},
				{"Team update", 0.0539732530010498},
				{"One-shot imitation learning", 0.05078921937077204},
				{"Proximal Policy Optimization", 0.0498038905999105},
				{"Moving AI governance forward", 0.04572699149265275},
				{"Meta-learning for wrestling", 0.0454000095983107},
				{"Evolved Policy Gradients", 0.04158881586178093},
				{"OpenAI Fellows Fall 2018: Final projects", 0.03943967314266656},
				{"OpenAI Scholars 2019: Meet our Scholars", 0.038789568640314916},
				{"OpenAI Scholars 2019: Final projects", 0.03778609184230498},
				{"Learning from human preferences", 0.03755160368370911},
				{"Machine Learning Engineer, Moderation", 0.03649691358024692},
				{"Team++", 0.03529850746268657},
				{"Team update", 0.033222124670763825},
				{"Implicit generation and generalization methods for energy-based models", 0.0318004571735915},
				{"Learning concepts with energy functions", 0.031469694617551176},
				{"Evolution through large models", 0.030041282946967292},
				{"Software Engineer, Security Data Platform", 0.02868753032508491},
				{"OpenAI supporters", 0.028581787419179407},
				{"OpenAI Five defeats Dota 2 world champions", 0.028201204364330146},
				{"Hierarchical text-conditional image generation with CLIP latents", 0.028195040534096327},
				{"Benchmarking safe exploration in deep reinforcement learning", 0.028014688462449654},
				{"Attacking machine learning with adversarial examples", 0.027230857800805988},
				{"Gathering human feedback", 0.027152698048220437},
				{"Third-person imitation learning", 0.026741293532338308},
				{"Sim-to-real transfer of robotic control with dynamics randomization", 0.02648969534050179},
				{"OpenAI customer story: Stripe", 0.025712111328549683},
				{"OpenAI partners with Scale to provide support for enterprises fine-tuning models", 0.025661892361111112},
				{"Report from the self-organizing conference", 0.02521321961620469},
				{"Team update", 0.024945941669743156},
				{"OpenAI Fellows Fall 2018", 0.024945941669743156},
				{"IT Network Engineer", 0.024604660840615897},
				{"Senior Manager, Strategic Sourcing (Non-Technology)", 0.024452026468155502},
				{"OpenAI Baselines: ACKTR & A2C", 0.02438584280669193},
				{"Better language models and their implications", 0.023659463785514204},
				{"One-shot imitation learning", 0.02363109512390088},
				{"Procgen Benchmark", 0.023351105845181675},
				{"Learning with opponent-learning awareness", 0.021539162112932605},
				{"Spinning Up in Deep RL: Workshop review", 0.021200304782394336},
				{"The power of continuous learning", 0.01999915436979409},
				{"Concrete AI safety problems", 0.019448213478064226},
				{"Microsoft invests in and partners with OpenAI to support us building beneficial AGI", 0.018775801841854557},
				{"New and improved content moderation tooling", 0.018401805166511048},
				{"Preparing for malicious uses of AI", 0.01833688699360341},
				{"WebGPT: Improving the factual accuracy of language models through web browsing", 0.01819790704832256},
				{"Solving Rubik’s Cube with a robot hand", 0.017803372478169228},
				{"Report from the OpenAI hackathon", 0.017782623406894998},
				{"DALL·E now available without waitlist", 0.017755255255255255},
				{"Frontier risk and preparedness", 0.016888031990859753},
				{"Research", 0.016611062335381913},
				{"Software Engineer, Distributed Systems", 0.01642361111111111},
				{"OpenAI Scholars 2020: Applications open", 0.016304160490848298},
				{"Learning a hierarchy", 0.01597217532248261},
				{"Competitive self-play", 0.01579351564326021},
				{"Improving language understanding with unsupervised learning", 0.015703851261620185},
				{"Research", 0.01545751633986928},
				{"OpenAI cybersecurity grant program", 0.014964565932675271},
				{"OpenAI’s API now available with no waitlist", 0.01486299648064354},
				{"Security Engineer, Detection and Response", 0.01482944569852019},
				{"Security Engineer, Detection & Response", 0.014598765432098766},
				{"Software Engineer, Privacy", 0.01437515195720885},
				{"Procgen and MineRL Competitions", 0.01421957671957672},
				{"Robots that learn", 0.014158285440613027},
				{"Research Engineer, Preparedness", 0.014007344231224829},
				{"Kernel Engineer", 0.01388888888888889},
				{"Generative models", 0.013690435990699381},
				{"National Security Threat Researcher", 0.013601334253508167},
				{"Data Engineer, Applied Engineering", 0.01354524627720504},
				{"OpenAI Supplier Code of Conduct", 0.013393362781741988},
				{"Frontier AI regulation: Managing emerging risks to public safety", 0.013244847670250895},
				{"Security Engineer, Partnerships", 0.013060525734481996},
				{"Introducing Superalignment", 0.012983091787439612},
				{"Research Engineer, Superalignment", 0.01280593458956032},
				{"Research Scientist, Superalignment", 0.012731481481481483},
				{"Generalizing from simulation", 0.012537107718405428},
				{"Roboschool", 0.01246573898376555},
				{"HW/SW Co-design Engineer", 0.012418609535811805},
				{"OpenAI Fellows Winter 2019 & Interns Summer 2019", 0.012407208246990006},
				{"OpenAI LP", 0.011277478422583568},
				{"OpenAI Baselines: DQN", 0.011277478422583568},
				{"March 20 ChatGPT outage: Here’s what happened", 0.011210656048540008},
				{"GPT-2: 1.5B release", 0.010858585858585859},
				{"Partnership with American Journalism Project to support local news", 0.0104775828460039},
				{"Manager, Enterprise Platform Sales", 0.010378269264525189},
				{"Our approach to AI safety", 0.009493416827231855},
				{"Quantifying generalization in reinforcement learning", 0.009200902583255524},
				{"CLIP: Connecting text and images", 0.008986928104575163},
				{"Frontier Model Forum", 0.008829898446833932},
				{"Multimodal neurons in artificial neural networks", 0.008782679738562092},
				{"How should AI systems behave, and who should decide?", 0.008678262145897548},
				{"Learning Day", 0.00855396412037037},
				{"ChatGPT plugins", 0.008261792267451508},
				{"OpenAI Five", 0.007879131130063966},
				{"Learning Montezuma’s Revenge from a single demonstration", 0.007582923192843516},
				{"Improving language model behavior by training on a curated dataset", 0.007568484383000512},
				{"Why responsible AI development needs cooperation on safety", 0.007423450570491392},
				{"Our approach to alignment research", 0.007307502162897046},
				{"AI safety via debate", 0.007218508683575986},
				{"Learning Montezuma’s Revenge from a single demonstration", 0.007056331304451605},
				{"AI safety via debate", 0.006717223358327654},
				{"Fine-tuning GPT-2 from human preferences", 0.0066040238470882265},
				{"DALL·E: Creating images from text", 0.006072861498956829},
				{"OpenAI Scholars 2019: Final projects", 0.005860342947764893},
				{"Democratic inputs to AI", 0.005553207476284399},
				{"Learning dexterity", 0.005230449398443029},
				{"Better language models and their implications", 0.0049437685521969985},
				{"Research index", 0.004923083328129229},
				{"Evolution strategies as a scalable alternative to reinforcement learning", 0.004642716921868865},
				{"Research index", 0.004581202541453587},
				{"OpenAI’s Approach to Frontier Risk", 0.004317860240083984},
				{"DALL·E 2 pre-training mitigations", 0.0038575716056632086},
				{"Scaling Kubernetes to 7,500 nodes", 0.0038194444444444443},
			},
		},
	}

	indx := indexInIt()
	indx.Open()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if urls, found := dlookup(tc.term, &indx, true); found {
				for idx, ti := range urls {
					if ti.url != tc.want[idx].url || ti.tfidf != tc.want[idx].tfidf {
						t.Errorf("%v failed\nGot: %v: %v Expected %v: %v\n", tc.name, ti.url, ti.tfidf, tc.want[idx].url, tc.want[idx].tfidf)
					}
				}
			}
		})
	}
	indx.Close()
}
