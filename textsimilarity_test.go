package textsimilarity

import (
	"reflect"
	"testing"
)

var testCorpus = []string{
	`berkshire hathaway investment boosts paytm’s valuation by 20% | the smarter way to get your business news - subscribe to bloombergquint on whatsappthe valuation of paytm’s parent one97 communication ltd. jumped 20 percent after it raised funds from billionaire investor warren buffett’s berkshire hathaway inc.paytm is valued at $10 billion (rs 70,634 crore) after the deal in which berkshire hathaway infused $300 million (rs 2,103 crore), documents filed by the registrar of companies and calculations by bloombergquint showed. berkshire hathaway’s arm bh international holdings bought 1,702,713 shares at$176.18 (rs 12,352.6) apiece, the value as of aug. 28 when the board approved the infusion.in may, one97 shares were worth rs 10,560 apiece, according to the filings, valuing the company at $8.34 billion (rs 58,586 crore).the investment gives berkshire hathaway a 2.9 percent stake in the company, whichcounts alibaba group as its largest shareholder with 38.4 percent stake, followed by softbank group corp and saif partners that hold close to 20 percent each.one97, which started in 2010 as a mobile recharge services provider, created a host of payment solutions such as a digital wallet—and the paytm payments bank is now india’s second most-valuable startup after flipkart.. read more on technology by bloombergquint.`,
	`next lifts profit guidance, plays down brexit threat | clothing retailer next raised its profit guidance after better-than-expected trading in late summer and said it was well prepared should britain crash out of the eu without a deal, sending its shares higher.`,
	`gold prices surge today, silver follows | gold prices jump by rs 175 to rs 31,725 per 10 gram, supported by positive global cues and continued buying by local jewellers`,
	`insight: how trump split mexico and canada in nafta talks | a day after winning the mexican presidential election, andres manuel lopez obrador took a congratulatory call from u.s. president donald trump. but trump had something more important on his mind: would mexico's new president consider a bilateral trade deal?`,
	`trump, brexit hit german growth, xenophobia also a threat | germany's bdi industry association on tuesday lowered its 2018 growth forecast and warned of a potential downturn, citing weaker demand for german exports due to u.s. trade policy and brexit, as well a threat to the economy from xenophobia at home.`,
	`sensex rallies 347 pts after seesaw trade; nifty50 tops 11,050 | volatility index india vix eased 5.30 per cent on tuesday to 16.5025.`,
	`rbi needs to step up pace of liquidity infusion, say analysts | the smarter way to get your business news - subscribe to bloombergquint on whatsappwithbanks borrowing more from the reserve bank of india than they have in at least the last two years, market participants feel that the central bank needsto intervene more decisively to calm nerves.the liquidity deficit has moved closer to rs 1.5 lakh crore in recent weeks driven by reasons ranging from advance tax outflows and increased demand during the festive season, to intervention in the currency markets. other factors such as the approaching election season, when demand for currency increases, could mean that liquidity tightness will persist.as such, most economists and money market analysts seethe need for more frequent bond purchases under the rbi’s open market operation program. on monday, the rbi announced bond purchases worth rs 10,000 crore - the second such purchase of the month.open market operations (omo) by the central bank can help contain the domino effects of liquidity tighteningin the financial markets, said bank of america merrill lynch in a report on tuesday. the rbi will need to inject $40 billion (about rs 2.8 lakh crore) by march 2019, the report estimates while adding that a calendar of bond purchases under the omo scheme would give markets more clarity on the outlook for liquidity.bond purchases by the rbi will help bring down yields and reduce foreign portfolios from the debt markets, which in turn could support the rupee, the report added. foreign investors have sold over rs 46,000 in the debt markets so far this year.kotak mahindra bank, in a report, estimated that the rbi would need to purchase bonds to the tune of rs 1.5-2 lakh crore by march 2019. economists at the bank believe that liquidity will remain tightin the second half due to increased currency demand during the festive season, intervention in the forex markets and cash build-up by the government.but even with continued bond purchases, kotak mahindra bank expects short term lending rates would continue to remain elevated. “it is unlikely to completely offset the liquidity tightness implying that the short-term rates would still likely remain elevated,” the report said.commenting on the troubles specifically in the nbfc sector, the kotak economists said they do not expect that to precipitate into a full fledged liquidity crisis. the recent squeezein money market liquidity has mostly been on the back of portfolio adjustments in debt funds, which had a relatively larger exposure to specific entities, said the report.should the rbi feel the need to address the specific liquidity needs of either mutual funds holding debt paper or non-bank lenders, it would need to craft out a special facility. one option could be to open a special liquidity window for mutual funds. it has done so at least twice inthe past decade in 2008 and 2013. the facilities, while barely used, send a comforting message to the market and market participants, said an analyst who spoke on condition of anonymity.channeling liquidity into nbfcs may prove to be more challenging if banks remain averse to lending to that segment. in such a situation, a line of refinance could be considered, this analyst added.. read more on business news by bloombergquint.`,
	`gold climbs ₹175 on global cues, jewellers’ buying | gold prices surged by ₹175 to ₹31,725 per 10 gram at the bullion market on tuesday, driven by a firm global trend amid persistent buying by local jew`,
	`indonesia's pertamina backs $100 billion spend to boost oil output | indonesia's state-owned energy company, pertamina, needs $100 billion over the next 12 years to boost oil output growth, a senior official said on tuesday.`,
	`indian hotel chain oyo to raise $1 billion from softbank, others | hotel chain oyo hotels said on tuesday it would raise $1 billion from existing investors, including japan's softbank group, to grow its business in india and china, and expand into new international markets.`,
	`f&o: nifty50 has to hold above 11,080 to bounce towards 11,171 | nifty has to cross and hold above 11,080 level to witness bounceback towards 11,171 level.`,
	`germany to hold diesel summit amid differing views on how to tackle crisis | there will be another high-level diesel summit in the german  chancellery on friday, two government sources said on tuesday, as the transport minister said his top priority was to ensure that diesel owners can swap their old vehicles for cleaner ones.`,
	`boj's kuroda highlights need to look at downside of easy policy | bank of japan governor haruhiko kuroda said on tuesday the central bank has entered aphase where it must consider not just the merits but the side-effects of its massive stimulus programme in a "balanced manner".`,
	`oyo hotels raises $1 billion, to fund expansion in china, uk, malaysia, nepal | financial services company j.p. morgan is acting as the exclusive financial advisor to oyo on this fundraising.`,
	`gold steady as investors await fed rate view | gold steadied on tuesday as the dollar held firm ahead of a u.s. federal reserve meeting, with the precious metal's upside capped by strong u.s. economic data that continues to underpin the greenback.`,
	`u.s. approval of $330 million military sale to taiwan draws china's ire | the u.s. state department has approved the sale to taiwan of spare parts for f-16 fighter planes and other military aircraft worth up to $330 million, prompting china to warn on tuesday that the move jeopardized sino-u.s. cooperation.`,
	`finance minister arun jaitley sees economy growing at 'around 8 percent' | finance minister arun jaitley said on tuesday he expects india's economy to sustain an annual growth rate of around 8 percent on the back of measures taken by the government such as a unified tax code and a new bankruptcy law.`,
	`ipo to meet capital adequacy and growth requirements for next 3-5 years: sushil kumar agarwal, aavas financiers | “our assets’ average tenure is 13 years. after prepayment, it is seven, eight years.”`,
	`german institutes to cut 2018 growth forecast to 1.7 percent: sources | germany's economic institutes will lower their 2018 growth forecast for europe's largest economy to 1.7 percent from their spring prediction of 2.2 percent, two government sources told reuters on tuesday.`,
	`jaitley asks psbs to take effective action against frauds, wilful defaults | finance minister arun jaitley tuesday asked public sector banks to take "effective action" in cases of fraud and wilful loan defaults.`,
	`reviewing the performance of state-owned lenders, he exuded confidence that formalisation of the economy would ...`,
	`hdil reaches settlement with bank of india over dues worth rs 5.11 bn | bank of india accepted hdil's structured payment plan to pay rs 5.1 billion in one year's time`,
	`nifty closes nearly 1 percent higher as financials recover | indian shares snapped a five-session losing streak and ended higher on tuesday, helped by gains in financials such as housing development finance corp ltd and hdfc bank.`,
	`india ratings downgrades debt of il&fs arm | rival rating agency icra had last month junked the ratings of most of the group companies.`,
	`novartis to cut 2,550 jobs in switzerland, uk in profit push | novartis will cut 2,550 jobs in switzerland and britain over four years, it said on tuesday, as the swiss drugmaker strives to boost profits and focus on new medicines.`,
	`nclat’s hyundai order: does the cci need to step up its game? | the smarter way to get your business news - subscribe to bloombergquint on whatsappthe competition commission of india had imposed a penalty of rs 87 crore on  hyundai motor india ltd., which was set aside by the national company law appellate tribunal last week.the tribunal’s observations may require the regulator to evaluate how it analyses evidence and articulates its views, according to experts.. read more on law & policy by bloombergquint.`,
	`bpcl, hpcl, ioc shares down over 2% as crude hits fresh 4-year high | jp morgan said us sanctions on iran could lead to a loss of 1.5 million bpd.`,
	`will ensure il&fs doesn't collapse, all options open: lic chief v k sharma | state-owned insurer lic tuesday said it will not allow debt-ridden il&fs to collapse and explore options to revive it.`,
	`the life insurance corporation (lic) has the largest shareholding in il&fs.`,
	`il&fs financial services, a group company of il&fs ...`,
	`click, post, get insta gains | from wine ­tasting to fitness, influencers are taking the social media platform by storm doing what they love and makingmoney in the age of influencer marketing.`,
	`oil boost and italy budget hopes support european shares; next shines | oil majors bp and shell rose more than 1 per cent after brent hit a fresh four-year high`,
	`don’t get greedy and compromise on quality: dhirendra kumar | “it is not a calamity because 1% to 3% decline take away 3-6 months of returns.”`,
	`canara bank drags gtl group to nclt on default of over rs 10 billion | as of march 2018, gtl had debt of rs 65.02 billion while that of gtl infrastructure had rs 49.56 billion`,
	`full-blown trade war would cost jobs, growth and stability - wto's azevedo | a full-blown trade war would have serious effects on global economic growth and there would be no winners of such a scenario, the director-general of the world trade organization (wto), roberto azevedo, said on tuesday.`,
	`new tariffs drive china shares lower; property firms drop on pre-sale system review | shares of chinese property developers plunged; real estate sub-index lost 4.4 per cent`,
	`instagram co-founders resign in latest facebook executive exit | instagram on monday said co-founders kevin systrom and mike krieger have resigned as chief executive officer and chief technical officer of the photo-sharing app owned by facebook inc, giving scant explanation for the move.`,
	`money problems? get newton’s help | newton’s laws of motion, which form the foundation of classical mechanics, can also be used to explain decision-making in personal finance from a behavioural perspective`,
	`cutting cash reserve ratio can improve liquidity: finance ministry official | cutting banks' cash reserve ratio (crr), or the amount of funds they set aside with the central bank, are among options that the reserve bank of india (rbi) could look at to improve liquidity in the system, a finance ministryofficial said on tuesday.`,
	`nikkei rises for 7th session as chip-related stocks offset machinery stocks' declines | the nikkei share average gained 0.3 per cent to 23,940.26`,
	`china says u.s. putting 'knife to its neck', hard to proceed on trade | a senior chinese official said on tuesday it is difficult to proceed with tradetalks with the united states while washington is putting "a knife to china's neck", a day after both sides heaped fresh tariffs on each other's goods.`,
	`flair writing industries files ipo papers with sebi | the initial public offering comprises equity shares of face value of rs 10 each of flair writing aggregating up to rs 450 crore.`,
	`why il&fs picked this route to solvency | the smarter way to get your business news - subscribe to bloombergquint on whatsappinsolvent infrastructure finance firm il&fs ltd. has sought to reorganise its financial affairs under a scheme of arrangement involving itself and 40 group companies. the companyhas filed an application for such a scheme with the national company law tribunal, it said in a statement today.the resolution of il&fs presents a complex challenge. not only because it involves the need for additional capital as well as debt restructuring. but also because unfortunately there is no specific legal framework to work with. yet.il&fs is a core investment company as per classification of the reserve bank of india. that is, a non-banking financial company whose core business is to make investments. in rbi parlance, it is better described as a cic - si - nd or a sytemically important non-deposit accepting cic.this classification is important as it determines what route il&fs can take to resolve its insolvency.besides the complex nomenclature, that il&fs has 24 direct subsidiaries, 132 indirect subsidiaries and six joint ventures adds layers to the resolution problem.no specific insolvency framework for financial firms, yethad the financial resolution and deposit insurance bill been passed, a specific framework for resolution of financial firms would have been available.because standard insolvency processes are not suited to financial firms of systemic importance, the bill sought to create “a credible resolution regime under an expert statutory institution that is able to ensure efficient, orderly and fair resolution of financial firms,” according to the committee that worked on its draft.that expert statutory institution was conceived as a resolution corporation that would count among its members representatives of the finance ministry and key economic regulators —rbi, sebi, irdai and pfrda— along with independent members.the resolution corporation would have powers to takeover any financial firm classified as having “critical” risk of failure, administer a resolution process within two years or a liquidation process if required.this special mechanism brought to bear the force of the government and the combined expertise of regulators when resolving a financial firm of systemic importance.but the frdi bill ran into other trouble and hence, unlike several other countries, no such special route is yet available for resolution of insolvent financial firms.also read: defaulting shadow lender faces india insolvency filingcould il&fs have taken the ibc route?the insolvency and bankruptcy code, 2016 does not cover a “financial service provider”.one way into the ibc though for a company like il&fs was by special notification of the government, under section 2 of the code. this route has not been used so far, but then the ibc itself isjust under two years in implementation.even if the ibc were to apply in such a case, the layered complexity of il&fs could have made for a tough, protracted resolution, experts say.also read: rbi calls for a meeting with il&fs shareholderscompany law solutionthe companies act, 2013 does provide for a solution, even if not an ideal one.section 230 of the law provides for a scheme of compromise or arrangement between a company and its creditor or shareholders. in simple terms, a scheme to reorganise the company’s financial structure.any debt restructuring under such a scheme would need approval of at least 75 percent of the secured creditors in value. or else, such a scheme would need similar approval from shareholders. and approval from the national company law tribunal.this route bears two obvious advantages—that equity rights are maintained, important in the il&fs context as major shareholders areexpected to bring in fresh funds. in an insolvent firm equity is worth nothing.and the scheme structure offers considerable flexibility in crafting a multi-faceted solution.though as one senior corporate lawyer pointed out, on condition of anonymity, schemes have not been tested in a group insolvency situation.and the lack of certain provisions may make them unwieldy.1. for instance there is no explicit provision of a moratorium or standstill. but earlier judgments, under the 1956 act, indicate that the courts can grant moratorium, the lawyer pointed out.2. unlike ibc, there are no fixed timelines ina scheme.3. also, a resolution plan approved by the nclt under the ibc has an overriding effect over all other laws, such as income tax act. however, schemes approved by the nclt under section 230 of the companies act, 2013 do not have any such benefit. they are subject to all other applicable laws andto that extent it may reduce the effectiveness of the scheme.other difficulties may include the lack of a comprehensive framework covering all creditors, like the committee of creditors in an insolvency resolution proceeding under ibc.but given the lack of immediate other options, il&fs has picked thisroute to nurse itself back to health. details of the scheme have not been disclosed so far except that it will involve il&fs and 40 other group companies. these include il&fs transportation networks ltd., il&fs engineering & construction company ltd., il&fs energy development company ltd., il&fs maritime infrastructure company ltd., il&fs environmental infrastructure & service ltd., il&fs township & urban assets ltd., hill country properties ltd. and many of their subsidiaries.the first step towards returning to solvency has been taken.shortly before il&fs announced the filing of a scheme, bloombergquint discussed the pros and cons of ibc vs companies act with mr umarji, former legal adviser to the indian banks’ association, and bhargavi zaveri, researcher at igidr.watch that discussion here.. read more on business news by bloombergquint.`,
	`germany's merkel says eu battery cell production 'extremely important' | it is "extremely important" that the european union develop its own battery cell production to secure its role in the automotive industry as it shifts to electric mobility, german chancellor angela merkel said.`,
	`flair writing industries files ipo papers with sebi | the initial public offering comprises equity shares of face value of rs 10 each of flair writing aggregating up to rs 450 crore, according to the draft red herring prospectus filed with sebi`,
	`not ruling out raising stake in il&fs: vk sharma, lic chairman | il&fs has defaulted on five of its obligations since august & put headquarters on the block.`,
	`ecb not planning to speed up normalisation, praet says | the euro zone economy remains on course for the european central bank to cut stimulus by one more step at the end of the year, but normalisation will not be any quicker than projected in june, ecb chief economist peter praet said on tuesday.`,
	`exclusive: china steel giant baowu in talks to take over rival magang - sources | top chinese steelmaker china baowu steel group is in talks to take over rival magang group, three sources familiar with the discussions said, a deal that would help entrench the nation's position as a serious competitor in global steel markets.`,
	`nbfc crisis: what went wrong & can it blow up in your face? | when il&fs first defaulted on its obligations, it made some analysts sit up and take note.`,
	`on the trail of a deadly disease that cuts down farmworkers in their prime | the smarter way to get your business news - subscribe to bloombergquint onwhatsapp(bloomberg businessweek) -- the tropical sun was beating down on homestead, a city on the southern tip of florida, when valerie mac, a nurse-scientist and occupational health specialist from emory university, pulled her van into the driveway of a small vegetable farm. its owner, isodoro gustavo, walked out to greet her, his arms caked with dirt and his face streaked with sweat and grime.mac and her two research assistants followed gustavo intohis home, where they sat him down to take his blood pressure, heart rate, and temperature and instructed him to provide a urine sample. they also drew his blood, placing the vial in a portable analysis machine. “we process it,” mac says, “pretty much like you would do in a hospital or a lab, and then we put it on dry ice right away to preserve it and get it back to the school to analyze it.”but on this particular day in august, with the mercury pushing past 90f, the machine kept overheating. the irony was not lost on mac and her crew, who were in south florida gathering data for a study about the wayextreme heat conditions affect the bodies of agricultural workers. more specifically, they’re looking for signs of stress on the kidneys, which can lead to a disease called chronic kidney disease of unknown origin, or ckdu.as its name indicates, the causes of ckdu are poorly understood. what’s not in question is that it’s deadly. symptoms—including vomiting, exhaustion, and weight loss—often don’t appear until the disease is well advanced, by which time damage to the kidneys cannot be reversed. without access to dialysis or a kidney transplant, there’s little hope of survival.ckdu first gained international recognition for its impact on sugar cane workers in nicaragua, where it’s killed at least 20,000 young men in the past decade, according to experts who describe it as an epidemic. the disease has also cropped up in other developing countries, including brazil, egypt, india, qatar, and sri lanka,predominantly among those such as field hands and construction workers who perform strenuous work outdoors. there’s no data on the number of cases worldwide.because it primarily afflicts poor people in poor countries, research around ckdu has been spotty. it’s well known that diabetes and high blood pressure can lead to chronic kidney disease, which affects 14 percent of the u.s. population. in contrast, there’s a dizzying number of hypotheses about what causes ckdu, including heat, dehydration, and exposure to pesticides or other toxins; there may also be a genetic component, according to some scientists.some hospitals in the u.s. are reporting a sharp increase in patients whose symptoms are consistent with the disease, which may lend new impetus to the search for answers.at the university of miami’s miller school of medicine, dr. alessia fornoni says the number of ckdu diagnoses has doubled sinceshe became chief of the nephrology department two years ago.“these patients show up with a picture that is very unique,” says dr. david sheikh-hamad, anephrologist affiliated with baylor st. luke’s medical center in houston. “they look very healthy, and if you didn’t check the blood, you wouldn’t haveknown that they had kidney failure.”while most experts agree that ckdu is a multifactor disease, there’s a growing consensus that sustained exposure toheat, along with dehydration, plays a leading role. this is where the emory team fits in. mac fell into ckdu research by accident after a previous study she conducted on a population of agricultural workers aged 18 to 54 in florida showed striking results: 33 percent of them experienced acute kidney injury (aki), a precursor to ckdu. the odds of a worker experiencing aki increased by almost half for each 5f increase in the heat index, a combined measure of temperature and humidity.mac’s current study, which involves about 100 people, is focused squarely on ckdu. cristina, a worker at a plant nursery near homestead, who requested that her last name be withheld, says she enrolled in the study as a way to keep tabs on her health for free. she’s acquainted with the dangers in her line of work. back in her hometown in northern mexico, her mother developed a chronic kidney condition after years of working outdoors. “i had realized that there are people in the country that are exposed to the sun,” says cristina, speaking in spanish. “there are a lot of problems that happen, and they don’t attend to them because, well, sincerely, the truth is that they don’t have the resources to get treated.”agricultural workers in this part of florida rarely have health insurance—many are undocumented immigrants from mexico and central america. (a 2018 analysis by thenonprofit national center for farmworker health determined that only 14 percent had access to employer-provided health care and just 9 percent were self-insured). as part of her study, mac helps organize training workshops for farm owners and their employees on how to guard against heat exhaustion.ruben and elude sanchez run a dragon fruit farm in redland, fla., and have a working relationship with mac and the emory study. “the biggest challenge here,as a medium-sized dragon fruit farm, is to find people who can work at the harvest time,” says ruben, speaking in spanish. “all of the farms and nurseries around here will tell you the same. hypothetically, if they were to be affected by a serious health condition, it’s going to be that much harder to get people to come to work.”mac says many producers, particularly those running larger-scale operations, don’t want their workers getting involved in studies such as hers because it can bring unwelcome scrutiny of working conditions and could stoke tensions between management and labor. “there’s mistrust on both sides,” says mac, who’s had to develop diplomatic skills in addition to her scientific training. “and for us, we’re like the third wheel.”in may, the national institute of diabetes and digestive and kidney diseases hosted a conference in bethesda, md., that many described as the first concreteattempt to establish a research agenda for ckdu. mac and the team from emory were among the 120 people in attendance. “we moved it outside the sphere of kidney docs talking about a kidney problem and talked about it in terms of an environment, kidney, and gene interaction,” says susan mendley, a program director at the institute. “we’ve changed the conversation, and i think that’s a real step.”while mendley hopes the attention generated by the conference will help unlock federal funding for ckdu research, she’s not certain that will happen. “the meeting was very effective,” she says, “but now can i turn that into an initiative that lets me access the cash register? not sure yet.”climate change lends urgency to the type of research mac is doing in florida, which, of the 50 u.s. states, is expected to see by far the greatest increase in the dangerous combination of heat and humidity over the next several decades, according to a 2016 report by climate central, an independent organization that brings together scientists and journalists. miami, an hour’s drive north of homestead, is projected to log at least 100 additional “dangerous heat days,” those when the heat index tops 104f, by 2050.“we’re not as hot as central america,” says mac, sitting in a spare room at the farmworker association of florida office in homestead, which she and her crew have transformed into a makeshift lab. “but the problem is so severe in central america, why would we wait until we see that here?”to contact the editor responsible for this story: cristina lindblad at mlindblad1@bloomberg.net©2018 bloomberg l.p.. read more on business news by bloombergquint.`,
	`yes bank shares slump 12% ahead of board meeting | in intraday trade, yes bank shares fell as much as 12.2% and touched a low of rs 198.65 per share onbse`,
	`dhfl stock back to old ways, plunges 20% | dhfl had rebounded 12 per cent in the previous sessions amid nbfc meltdown.`,
	`tata steel is not willing to give even an inch on the acquisitions front | at a time when its balance sheet is getting stretched by acquisitions, with more in the pipeline, investors would like to see how tata steel intends to make this acquisition work`,
	`oyo raises $1 bn from investors led by softbank; firm now valued at $5 bn | this e-series funding round makes oyo the most valued hospitality company in the country, ahead of tata group's indian hotels company and eih`,
	`ola to monitor rides real time for passenger safety | in a bid to step up passenger safety, indian cab-hailing major ola has rolled out a real-time ride monitoring system, the company said on tuesday.“ol`,
	`calm is returning to the market: mahendra jajoo, ‎mirae asset global investments | “the equity correction needs to be delinked with the panic in the debt market.”`,
	`oil hits four-year high as opec, russia resist output rise to offset iran sanctions | singapore (reuters) - brent crude oil prices hit a fresh four-year high on tuesday amid looming u.s. sanctions against iran and an apparent reluctance by opec and russia to raise output to offset the expected to hit to supply.`,
	`bnp paribas cardiff may trim stake in sbi life; stock hits new low | the stock slipped 6% to rs 553 on the bse in intra-day trade after the company said that bnp paribas cardiff may trim stake in the company to comply with minimum shareholding norms.`,
	`bad news for creditors as il&fs, 40 arms move nclt seeking debt protection | according to corporate law experts, by moving the mumbai branch of nclt, il&fs is seeking more time to repay its debt`,
	`hcl technologies hits record high; stock up 24% in three months | hcl technologies hit a record high of rs 1,125, up 2.6% in intra-day trade, surpassing its previous high of rs 1,107 recorded on april 23, 2018, on the bse.`,
	`omcs, aviation stks drop as crude rises; indigo, jet airways hit 52-wk lows | shares of hindustan petroleum (hpcl) was trading two and a half per cent lower at rs 242 while those of indian oil was down 1 per cent at rs 153.5.`,
	`falling rupee has a silver lining: rising software exports | the share of services exports, which mainly comprise software, has climbed to 7.3 percent of gross domestic product in june from 6.8 percent in march 2017`,
	`as il&fs crisis deepens, here are top 10 developments in the unfolding saga | the focus now is on the reserve bank of india's (rbi) summons to il&fs shareholders for a meeting on friday`,
	`india's oil demand to climb to 500 million tonnes per year by 2040: indian oil | india's crude oil demand is forecast to grow to 500 million tonnes peryear by 2040, but persistent increases in oil prices might act as a dampener for the rate of growth, partha ghosh, an executive director at indian oil corp said on tuesday.`,
	`oil just below four-year high as producers resist output rise to offset iran sanctions | oil prices on tuesday were holding just below four-year highs hit in the previous session, as looming u.s. sanctions against iran and unwillingness by the organization of the petroleum exporting countries (opec) toraise output supported the market.`,
	`after whatsapp, instagram founders quit amid tensions with zuckerberg | instagram, which now has more than one billion users, is a key driver of revenue for facebook`,
	`candidates can't be barred from polls on framing of criminal charges: sc | each candidate to declare criminal antecedents to ec before contesting an election, sc said`,
	`no more free food: india's 'cheap' airfares force airlines to change | jet airways' decision to make passengers pay for meals signals that full-serviceairlines will struggle in india in the near future`,
	`instagram co-founders resign in latest facebook executive exit | instagram on monday said co-founders kevin systrom and mike krieger have resigned as chief executive officer and chief technical officer of the photo-sharing app owned by facebook inc, giving scant explanation for the move.`,
	`global markets: stocks struggle as u.s.-china trade row revives growth concerns; oil elevated | asia stocks struggled on tuesday as a fresh round of u.s.-china tariffs and a surge in oil prices to near four-year highs added to worries about risks to global growth.`,
	`gold rate today: gold, silver trade lower ahead of us fed meet | the probability is 100 per cent of fed raising the interest rate in the fomc meeting.`,
	`the benchmark story: how equity indices hoodwinked desi fund managers | the best performing sensex stock was up about 65% and the worst was down about 30%.`,
	`sensex falls over 100 points on foreign fund outflow, rising oil prices | benchmark indices turned volatile after opening higher tuesday on foreign fund outflows, surging global crude oil prices, weak rupee and negative gl`,
	`yes bank gains 6% ahead of board meeting today | the board will decide the future course of action after the rbi curtailed the term of its md and ceo rana kapoor and asked the bank to look for a replacement by january 2019.`,
	`china says hard to proceed on trade with u.s. putting 'knife to its neck' | a senior chinese official said on tuesday that it is difficult to proceed with trade talks with the united states while washington is putting "a knife to china's neck", a day after both sides heaped fresh tariffs on each other's goods.`,
	`share market update: pharma stocks rise up to 4%; aurobindo pharma, sun pharma, lupin top gainers | the nifty pharma index was trading 1.68 per cent upat 10,163 at 10:35 am, with all components in the green zone.`,
	`il&fs seeks nclt relief for arrangement with creditors to 40 group firms | in a statement on tuesday, the company informed stock exchanges of the application it had filed with nclt`,
	`yes bank stock gains traction ahead of board meet | the scrip was up 1.75 per cent at rs 230.20 at around 9.40 am.`,
	`rafale contract or not, hal has its own issues | high inventories and stretched working capital cycle are weighing on the company’s earnings quality`,
	`instagram co-founders resign from social media company | kevin systrom and mike krieger founded the photo-sharing app in 2010 and sold it to facebook in 2012 for about $1 billion.`,
	`buzzing stocks: dhfl, rcom, gmr infra, yes bank | the counter of dhfl was buzzing as the most active stock in value terms too.`,
	`market falling? you could buy selectively in these 3 spaces: deven choksey | “not very sure of  a currency play in it and pharma sectors.”`,
	`beware! macd charts show these 106 stocks all set to take a tumble | the benchmark indices have fallen below the 100-day moving averages.`,
	`what are derivatives and how do they work | derivatives allow investors to bet on the possible future price of an asset, which can be anything from stocks, bonds and commodities to currencies and interest rates`,
	`opinion | investing in real estate funds can help diversify risk | managed real estate funds can invest in multiple projects across geographies`,
	`gold set to soar above $1,300/ounce, says boa | goldman sachs has also joined the chorus of bulls, seeing gold at $1,325 in 12 months.`,
	`stock market update: nbfc stocks trade mixed; dhfl, indiabulls housing finance down | weak global cues, rising crude oil prices and rupee's bumpy ride weighed on domestic equity benchmarks on tuesday.`,
	`electrosteel steels shares go haywire on delisting announcement | earlier vedanta had acquired management control of debt-laden electrosteel steels.`,
	`aavas financiers' rs 1,734-crore ipo opens: will you take the bet? | the price band for the public issue has been fixed at rs 818-821.`,
	`only a few companies provide health insurance to senior citizens | health insurance covers for parents, yourself and your children should be periodically reviewed`,
	`auto index hits 52-wk low; maruti suzuki, escorts down over 15% in 1-month | escorts, force motors, maharashtra scooters, vst tillers, eicher motors, atul auto, amara raja batteries, exide ind, motherson sumi systems and minda ind were down 10% to 28% in past one month.`,
	`senior china diplomat says confrontation with u.s. lose-lose | confrontation between china and the united states means both sides lose, and talks with washington cannot take place under threats and pressure, the chinese government's top diplomat state councillor wang yi told u.s. business leaders.`,
	`neepco bond issue gets only a single bid | neepco received only one bid for rs 13 crore with a 9.25 per cent yield (half-yearly).`,
	`share market update: midcaps, smallcaps trail sensex; berger paints, dhfl fall | bse midcap index was trading 0.76 per cent down at 15,106, while the bse smallcap index was trading 1.27 per cent down at 15,139.`,
	`interest income from fcnr exempt for nris, rnors | for those who are indian residents for tax purposes, interest income earned from fcnr deposits attracts tax`,
	`increasing tariffs may not be an effective solution to curb the rupee’s weakness | some other options that the government may have to choose from to deal with the rupee’s depreciation are increasing interest rates or opting for foreign bonds/nri bonds to shore up the country’s forex reserves`,
	`4 benefits of adding spouse as co-owner when buying a house | buying a house jointly with your spouse can save stamp duty as well as provide tax deduction`,
	`rupee falls 33 paise to 72.96 against us dollar | on monday, the local currency had ended sharply lower by 43 paise to 72.63 against the us dollar after crude prices soared ahead of impending us sanctions on iran.`,
	`stock market update: telecom stocks suffer as dot set to conduct audits; airtel, vodafone idea among top losers | the bse telecom index was trading 2.55 per cent down at 1,054 around 09:35 am.`,
	`indian equities not attractive enough to lure foreign investors back, bofaml says | the smarter way to get your business news - subscribe to bloombergquint on whatsappthe valuations of indian equities aren’t compelling enough to attract foreign investors.that’s the word from sanjay mookim, india equitystrategist at bank of america merrill lynch. “indian equities have overshot and are trading at all-time highs compared with other emerging market economies,” mookim said in an interaction with bloombergquint. he is cautious on the equity markets due to domestic risks with the elections nearing. there could be further downside in indian equities despite the recent correction, he said.“the domestic risk off is very likely in india with the elections coming due to which we are cautious on equity markets for a while,” he said.mookim doesn’t see valuations in his favour and would be cautious in chasing risks."the valuations are an indicator of risks. if the valuations are rising, one needs to derisk their portfolio, which means, the investor would avoid liquidity risks and not buy small- and mid-cap stocks unlike before." -  sanjay mookim, india equity strategist at bofaml.the rupee may have depreciatedagainst the u.s. dollar due to rising crude oil prices and a widening current account deficit, but mookim said that “we’re not in a situation to start defending the rupee aggressively”.“the rupee needs to keep up with its emerging market peers,” he said. “the currency depends on the government managingvolatility of the currency.”. read more on markets by bloombergquint.`,
	`sensex declines 100 pts, nifty50 below 10,950; telecom stocks fall up to 3% | sensex and nifty opened in the green but failed to hold altitude.`,
	`gold steady ahead of u.s. fed meeting; trade worries persist | gold held steady on tuesday as the dollar stood firm ahead of the two-day u.s. federal reserve meeting beginning later in the day, while simmering u.s.-china trade tensions kept investors nervous about risks to global growth.`,
	`over 50% of gen-next hope to manage family business | it is very important to have faith in the younger generation who is fully equipped to handle their exposure to world class business environment and practices which help them to perform better in real business situations, say experts`,
	`rupee near all time low against dollar on surging crude prices | at 9.15am, the rupee was trading at 72.95 to a dollar, down 0.44% from its friday’s close of 72.63. the home currency opened at 72.96 per dollar`,
	`bank fd vs stocks vs gold: how much you would have earned | here’s a look at how four commonly used asset types—equity, cash, gold and fixed income—have done in different periods`,
	`how concentrated is your mutual fund scheme? | here is a list of funds with the highest concentration of top 10 holdings`,
	`china to set up local government debt monitoring system - china daily | china is building a nationwide system to monitor the income and expenditure of local governments in a bid to control debt, the official china daily reported on tuesday, citing finance ministry officials.`,
	`is india’s soulless rally losing heart? | investors may manage to find some so-called safe havens now and then, but valuations will eventually need to catch up with fundamentals for all stocks`,
	`google ceo to meet us lawmakers, denies efforts to tweak search results | google denies it makes content decisions based on politics`,
	`rates of small savings schemes hiked: what should investors do? | interest rates on various small savings schemes, including ppf, senior citizen savings scheme (scss) and sukanya samriddhi scheme (sss) have been increased by 40 basis points (bps) for the october-december quarter.`,
	`u.s., japan push back trade talks to tuesday | japan and the united states will begin a second round of trade talks in new york on tuesday, japan's topgovernment spokesman said, amid concerns in tokyo that japan will face greater pressure to reduce its large trade surplus.`,
	`china says trade war to 'certainly' hurt us exporters, create opportunities to others | china is being forced to retaliate against the united states intheir trade dispute, and u.s. exporters including suppliers of liquefied natural gas would "certainly" be hurt, said chinese vice commerce minister wang shouwen.`,
	`after mutual funds, paytm money wants to start share trading now | if paytm wishes to apply for a stockbroker licence, bse would be delighted to welcome them, says bse ceo ashishkumar chauhan`,
	`markets live: indices turn positive, sensex up 150 pts; yes bank gains 3% | catch all live market action here`,
	`all you need to know going into trade on sept. 25 | the smarter way to get your business news - subscribe to bloombergquint on whatsappasian stocks were off to a muted start tuesday after u.s. stocks slipped on growing concern about the outlook for global trade and american politics.equity benchmarks were little changed in japan, which returned from a holiday, and australia. china’s markets also will resume trading after long weekend, while markets inhong kong and south korea are shut tuesday.the singapore-traded sgx nifty, an early indicator of nse nifty 50 index’s performance in india, was little changed at 10,996 as of 7:31 a.m.short on time? well, then listen to this podcast for a quick summary of the article!bq live`,
	``,
	``,
	`u.s. market checkstocks slipped on growing concern about the outlook for global trade and u.s. politics.the yield on 10-year treasuries was steady at 3.09 percent, near the highest since may.#bqmarketwrap | u.s. stocks fall on politics and trade risks; oil gains.read: https://t.co/q6szka8sxu pic.twitter.com/kn7pdbdyca— bloombergquint (@bloombergquint) september 25, 2018`,
	``,
	`also read: new u.s.-south korea pact spurs hopes for nafta, china deals. read more on markets by bloombergquint.`,
	`5 days, 16 judgments: a look into key cases to be delivered by cji misra | a glimpse into the cases shows that many of them are of profound significance for the future of the country`,
	`denying green cards to immigrants on benefits will only cost us more later | trump administration's clampdown on green cards for immigrants on public assistance 'cruel'`,
	`stocks pressured as u.s.-china trade fight revives growth fears; oil elevated | asia stocks struggled on tuesday as the latest round of u.s.-china tariffs revived fears the trade dispute would knock global growth, while crude oil was elevated near four-year highs after saudi arabia and russia ruled outimmediate production increases.`,
	`'search' for the future: google brings ai, more visuals to its site | the search engine will let users create collections of online content, and suggest related material that might be of interest`,
	`aavas financiers ipo opens today. should you subscribe? | the ipo comes amid choppy markets, which have been on a downward spiral over the past few trading sessions on the back of liquidity concerns regarding non-bank finance companies (nbfcs).`,
	`even as india's economy grows, its poor have hardly any upward mobility | for one group mobility is actually declining: indian muslims`,
	`oil prices surge as saudis, russia won't open spigots | global benchmark brent crude jumped more than 3 percent on monday to a four-year high above $80a barrel after saudi arabia and russia ruled out any immediate increase in production despite calls by u.s. president donald trump for action to raise global supply.`,
	`tata-usha martin deal: give more clarity on end use of funds, say jhawars | tata steel had signed a definitive agreement with usha martin on saturday for the acquisition of its steel business`,
	`special report: high-nicotine e-cigarettes flood market despite fda rule | the sleek juul electronic cigarettes have become a phenomenon at u.s. high schools, vexing educators and drawing regulatory scrutiny over their sweet flavours and high nicotine content.`,
	`wall street falls as u.s.-china tariffs kick in | the s&p 500 and the dow closed lower on monday after a new round of u.s.-china trade tariffs kicked in, dampening last week's hopes for talks between the two countries, and as investors awaited a widely expected interest rate hike by the federal reserve.`,
	`insight: the night a chinese billionaire was accused of rape in minnesota | with the chinese billionaire richard liu at her minneapolis area apartment,a 21-year-old university of minnesota student sent a wechat message to a friend in the middle of the night. she wrote that liu had forced her to have sex with him.`,
	`crude oil imports may be cut as price forecast put at $100 a barrel | benchmark brent crude oil futures surged 2% on monday to over $80 a barrel`,
	`trump calls new u.s.-south korea trade deal a historic milestone | u.s. president donald trump and south korean president moon jae-in signed a free trade agreement on monday that trump hailed as a "historic milestone in trade."`,
	`amazon made two moves for deliveroo, one nine months ago: telegraph | amazon.com inc  made two preliminary approaches for british online food delivery company deliveroo, the latest one about nine months ago, the telegraph reported on monday.`,
	`starbucks plans changes to company structure, layoffs | starbucks corp  is planning an organizational restructuring including leadership changes, according to a memo from its chief executive officer kevin johnson.`,
	`wall street's s&p communication sector starts with slight gain | the new s&p 500 communication services sector , which includes such high-profile namesas facebook inc , alphabet inc  and netflix inc , made its debut on monday with a small gain, after the largest-ever overhaul of wall street's broad business sectors.`,
	`stock markets fall on trade war pessimism; oil rallies | stock markets around the world retreated on monday amid concerns over the potential wider impact of a trade spat between china and the united states, while oil prices rallied to a four-year high after opec ignored u.s. calls to raise supply.`,
	`yes bank board to meet today after rbi directive on rana kapoor's tenure | yes bank to decide the future course of action after the rbi curtailed the term of its founding ceo rana kapoor`,
	`sebi to review dhfl, yes bank stocks over 'irregularities' in trading | sebi is also ascertaining the links between the sudden downgrading and rumours about credit default in various financial institutions`,
	`rbi to conduct open market operations to infuse liquidity of rs 100 billion | the purchase will happen through multi-security auction using the multiple price method`,
	`sensex, nifty set to experience a more painful bear hug | a chilling combination of reasons—past defaults, present insecurities and uncertain future expectations—has markets spooked and it will require more than empty assurances to restore confidence`,
	`odisha govt, centre spar over the issue of land allotment to iim-sambalpur | in june this year, the odisha govt, through a cabinet approval, sanctionedthe lease of government land measuring 181 acres for establishment of a permanent campus of iim-s`,
	`aavas financiers ipo opens today amid nbfc mayhem | aavas has fixed its price band at ₹ 818-821, aiming to raise up to ₹1,734 crore from the initial public offering (ipo)`,
	`stocks to watch: infosys, yes bank, dena bank, sbi life insurance | the smarter way to get your business news - subscribe to bloombergquint on whatsappasian stocks were off to a muted start tuesday after u.s. stocks slipped on growing concern about the outlook for global trade and american politics.equity benchmarks were little changed in japan, which returned from a holiday, and australia. china’s markets also will resume trading after long weekend, while markets in hong kong and south korea are shut tuesday.the singapore-traded sgx nifty, an early indicator of nse nifty 50 index’s performance in india, traded 0.1 percent higher at 11,002.50 as of 7:25 a.m.short on time? well, then listen to this podcast for a quick summary before the opening bell.. read more on markets by bloombergquint.`,
}

func TestSimilarity(t *testing.T) {

	ts := New(testCorpus)

	t.Run("similarity", func(t *testing.T) {
		cases := []struct {
			docA, docB string
			result     float64
		}{
			{
				docA:   "BSE, NSE to foray into commodity derivatives from Oct 1, to start with gold",
				docB:   "BSE, NSE to foray into commodity derivatives from Oct 1, to start with gold",
				result: 1.0,
			},
			{
				docA:   "Gold prices edge up as easing trade concerns hurt dollar",
				docB:   "Global gold prices edge up as easing trade concerns hurt dollar",
				result: 0.9,
			},
		}

		for _, tc := range cases {
			result, _ := ts.Similarity(tc.docA, tc.docB)
			if result < tc.result {
				t.Errorf("Similarity(%v,%v) did not return %v, got = %v", tc.docA, tc.docB, tc.result, result)
			}
		}

	})

}

func TestKeywords(t *testing.T) {

	ts := New(testCorpus)

	t.Run("keyword", func(t *testing.T) {
		cases := []struct {
			thresh []float64
		}{
			{
				thresh: []float64{0.2, 0.5},
			},
		}

		for _, tc := range cases {
			result := ts.Keywords(tc.thresh[0], tc.thresh[1])
			if len(result) == 0 {
				t.Errorf("Keywords(%v %v) did not return result, got %v", tc.thresh[0], tc.thresh[1], result)
			}
		}

	})

}

func TestTokenize(t *testing.T) {
	type args struct {
		s   string
		opt Option
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestSimple",
			args: args{
				s: "Hello to the World 123",
			},
			want: []string{"hello", "world"},
		},
		{
			name: "TestReplaceStopWords",
			args: args{
				s: "Hello to the World 123",
				opt: WithCustomStopwords([][]byte{
					[]byte(`hello`),
					[]byte(`world`),
				}),
			},
			want: []string{"to", "the"},
		},
		{
			name: "TestExtraStopWords",
			args: args{
				s: "Hello to the World 123",
				opt: WithExtraStopwords([][]byte{
					[]byte(`hello`),
				}),
			},
			want: []string{"world"},
		},
		{
			name: "TestWithBiGrams",
			args: args{
				s:   "Hello to the World 123 I am testing BiGrams one two three",
				opt: WithBiGrams(),
			},
			want: []string{"hello", "world", "testing", "bigrams", "one", "two", "three", "hello world", "world testing", "testing bigrams", "bigrams one", "one two", "two three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.opt != nil {
				ts := New(testCorpus, tt.args.opt)
				if got := ts.Tokenize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Tokenize() = %v, want %v", got, tt.want)
				}
			} else {
				ts := New(testCorpus)
				if got := ts.Tokenize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Tokenize() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
