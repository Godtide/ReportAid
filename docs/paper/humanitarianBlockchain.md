_Article_

Steve Huckle <sup>1*</sup>

1 University of Sussex, Sussex House, Falmer, Brighton, BN1 9RH, United Kingdom

\* Correspondence: s.huckle@sussex.ac.uk; Tel.: +44 (0)1273 606755

Received: (date); Accepted: (date); Published: (date)

## Abstract

This article describes a blockchain-based application for increasing the transparency of humanitarian financing. At the 2016 World Humanitarian Summit in Istanbul, humanitarian organisations made a 'Grand Bargain' to publish timely and transparent information about their funding - the summit proposed meeting the requirements of that bargain through a financial tracking system that satisfied the '3Ts' of transparency, namely traceability, totality and timeliness. Aid organisations have been investigating direct delivery of aid through blockchains, but there appears very little research, if any, into the use of blockchains to improve the transparency of their aid reporting. This paper fills that gap when discussing a prototype application that uses blockchains for aid reporting. It argues that blockchains include mechanisms ideally suited to meet the Grand Bargain's '3Ts', while adding a fourth - trust. This paper describes the prototype by showing how it might have been employed to report on the European Commission's 2015 response to the Ebola crisis in West Africa. The hope is that, ultimately, this paper, and the software it describes, leads to the deployment of a blockchain-based humanitarian aid reporting system that helps the UN Office for the Coordination of Humanitarian Affairs increase the transparency of humanitarian aid.

## Introduction

First, this paper provides some background to the application. Described next is the design of [ReportAid](https://github.com/glowkeeper/ReportAid), after which, an example use of the application is given via an imagined scenario where it is used to document the European Commission's 2015 response to the Ebola crisis in West Africa. The paper then analyses that scenario, before describing the application's limitations and proposals for future work.

## Background

The "Grand Bargain" (GB) made at the 2016 World Health Summit (WHS), committed to enhancing the transparency of mutual aid reporting [@WorldHumanitarianSummit_CommitmentsActionWorld_2016].

The amount of humanitarian financing contributed by the U.K. is a result of the official development assistance target of 0.7% of gross national income [@GOV.UK_OfficialDevelopmentAssistance_2018]. In 2015, that amounted to the U.K. spending £12.1bn [@Morris_RealityCheckHow_2017], a degree of funding that has received much criticism. For example, while giving evidence to a U.S. Senate Committee on foreign relations, ex-UK Prime Minister David Cameron suggested that much of the money goes to corrupt regimes, "If what we do is just have continued programs for countries that sometimes fail year after year after year, we just keep going, maybe that's not a good use of our money" [@MailOnline_StripAidMoney_2018].

Their have been efforts to address such criticisms. The 2016 World Humanitarian Summit (WHS) described five 'commitments to action', which outline the key responsibilities for aid:

1. Uphold the norms that safeguard humanity by enhancing compliance and accountability to international law.
2. Implement a new approach to forced displacement so that no one is left behind.
3. Achieve gender equality and greater inclusivity.
4. Instead of replacing local systems, reinforce them so that, eventually, there is no need for aid.
5. Invest in humanity by diversifying the resource base and increasing cost-efficiency [@WorldHumanitarianSummit_CommitmentsActionWorld_2016].

The WHS described those five key responsibilities as an 'Agenda for Humanity' [^2], whose aim was to transform the lives of 130 million people living in crisis-affected areas around the world  @WorldHumanitarianSummit_CommitmentsActionWorld_2016]. As a result, large humanitarian donors and aid organisations agreed to a "Grand Bargain" (GB) to improve the lives of people living in fragile situations because of crises [@WorldHumanitarianSummit_CommitmentsActionWorld_2016]. The GB made many commitments, including enhancing the transparency of mutual aid reporting, thereby strengthening accountability, helping decision-making and ultimately, improving the effectiveness of humanitarian efforts. The Inter-Agency Standing Committee (IASC), a forum that was founded by UN and non-UN humanitarian partners in 1992 to strengthen mutual assistance, formed a 'workstream' to carry out the GB's transparency commitment [^3]. The workstream's baseline report acknowledged that, while more information on aid funding had become available, there was still a need for more organisations to produce better quality humanitarian data [@DevelopmentInitiatives_BaselineReportImplementing_2017].

Consequently, the International Aid Transparency Initiative (IATI) was adopted as the United Nation's standard open-data format for documenting aid finance [@DevelopmentInitiatives_BaselineReportImplementing_2017]. Established in 2008, IATI became part of the movement for improving the reporting of charitable actions. Then, in 2013, various bodies of the United Nations (UN) began hosting IATI [@DevelopmentInitiatives_ImprovingHumanitarianTransparency_2017], after which it became the international framework for publishing open data on development cooperation and humanitarian assistance [@DevelopmentInitiatives_BaselineReportImplementing_2017]. Indeed, as of January 2017, over 500 humanitarian organisations are using IATI, including the governments of Japan, Sweden, the U.K. and the U.S., the European Commission's Humanitarian Aid and Civil Protection department, Oxfam, Save the Children, UNICEF and the World Food Programme [@DevelopmentInitiatives_ImprovingHumanitarianTransparency_2017].

The IASC's GB workstream identified the Financial Tracking Service (FTS) of the U.N. Office for the Coordination of Humanitarian Affairs (OCHA), as the humanitarian reporting platform where IATI data from different organisations, and various platforms, could be amalgamated and published for analysis by global actors [@DevelopmentInitiatives_ImprovingHumanitarianTransparency_2017]. The FTS came into being in 1992 as a result of a set of guiding principles for strengthening the coordination of humanitarian emergency assistance under U.N. General Assembly Resolution 46/182 [^4]. As of the end of 2016, three hundred and fifty humanitarian organisations reported financial information to FTS, including all significant government donors, all U.N. humanitarian agencies, Red Cross organisations, as well as 250 NGOs and private organisations [@DevelopmentInitiatives_ImprovingHumanitarianTransparency_2017]. FTS can import the latest versions of the IATI standard; thus IATI complements FTS by providing the technical publishing framework by which FTS can make available structured reports on aid efforts.

This paper examines whether blockchains have capabilities that can enhance the FTS and add further succour to the transparency initiative of the WHS. It does so through [ReportAid](https://github.com/glowkeeper/ReportAid), an appliucation that implements IATA on the blockchain.

## Application Design

The IATI open data standard defines specific entities that need recording. Figure 1, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) as a blockchain-based application that allows users to input, amend, and read IATA standard organisation and activity records. Hence, [ReportAid](https://github.com/glowkeeper/ReportAid) is a blockchain-based proof of concept that represents an implementation of OCHA's FTS.

![Figure 1: A Use Case Diagram for ReportAid](images/reportaidUseCaseDiagram.png)

Figure 2 shows [ReportAid](https://github.com/glowkeeper/ReportAid) implementing the IATI organisations standard via blockchain smart contracts. That standard is used to describe planned future budget information for aid funding [^5]. Described is a top-level IATIOrganisations element (this contains information such as the generation date of the report), which has at least one IATIOrganisation element (containing the report's language and currency type). That IATIOrganisation element links to other information describing the aid, such as the reporting organisation, associated country budgets and any supporting documentation.

![Figure 2: Organisations Smart Contracts](images/organisationsClassDiagram.png)

Figure 3 shows [ReportAid](https://github.com/glowkeeper/ReportAid)  implementing the IATI activities standard via blockchain smart contracts. A large number of fields describe IATI activities, and [ReportAid](https://github.com/glowkeeper/ReportAid), at the time of writing, has not implemented all of those; however, it does support all the mandatory fields, as well as one or two that are recommended [^6]. Figure 3 shows a top-level IATIActivies element (which contains information such as the generation date of the report). That features at least one IATIActivity element (containing information such as the default currency type and the degree to which the activity relates to humanitarian aid), which links to information such as the organisation participating in the activity, the sector and territory to which the activity belongs, as well as budgetary data. A single reporting organisation produces an activity report.

![Figure 3: Activities Smart Contracts](images/reportaidActivitiesClassDiagram.png)

## A Humanitarian Aid Report

The baseline report of the IASC GB transparency workstream cites the recent Ebola crises as an example of incomplete, inaccurate, inconsistent and often inaccessible information that impacted the humanitarian effort in the diseases' epicentre in West-Africa [@DevelopmentInitiatives_ImprovingHumanitarianTransparency_2017]. Indeed, the report says that the lack of an adequately planned response hindered attempts to alleviate the outbreak because, at the time, no donor, government or aid agency was able to gain an overarching overview of available resources. The suggestion is that, were all the Ebola humanitarian aid efforts documented using the IATI standard and published to the UN's FTS, organisations would have had a more accurate picture of what was needed, thus improving the effectiveness of their response.

Despite such shortcomings, many of the organisations that are already using IATI make information about their development contributions available through existing aid information portals; a widely-used example is an open-source platform called [d-portal](http://d-portal.org/) [^7]. That includes several IATI activities related to the world's response to the Ebola outbreak in West Africa. Hence, it contains data that [ReportAid](https://github.com/glowkeeper/ReportAid) can import, thereby demonstrating its suitability as a proof of concept for blockchain-based humanitarian aid reporting. One such activity documented on [d-portal](http://d-portal.org) was that carried out by the European Commission's Directorate-General for International Cooperation and Development (DEVCO), titled "A West African Response to Ebola" (AWARE), with IATI activity identifier XI-IATI-EC_DEVCO-2014/37785/0. That had a planned start date of 15th December 2015 and planned end date of 27th November 2018. The outgoing commitments of the activity totalled US$39,957,400. Its overall objective was to mitigate the harmful effects of the diseases' outbreak. Additionally, it was to contribute to the recovery of the most affected countries. Its specific purpose was to increase awareness of the diseases' symptoms and strengthen the resilience of primary healthcare systems. [Appendix A](./appendixAWAREXML.md) includes the XML describing AWARE.

Figure 4, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) creating the organisation record for AWARE.  

![Figure 4: Creating the organisation record for DEVCO](images/devcoOrgRecord.png)

The smart contract implementation of IATI Activities depends on a single top-level activities record. [ReportAid](https://github.com/glowkeeper/ReportAid) must create that before it can record any specific activity. Figure 7.5, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) using MetaMask to sign the transaction creating the necessary top-level activities record [^705].

![Figure 5: Creating the overarching activities record for the DEVCO Ebola activity](images/devcoActivitiesRecord.png)

Now the specific activity can be recorded. Figure 7.6, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) signing the transaction creating the activity relating to AWARE.

![Figure 6: Creating the activity record for DEVCO's AWARE activity](images/devcoActivityRecord.png)

Subsequently, anyone can read that activity record. Furthermore, that data can be trusted because the transaction that created that record was digitally signed. Figure 7, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) retrieving the information created in Figure 7.6.

![Figure 7: Reading the activity record for DEVCO's AWARE activity](images/devcoReadActivityRecord.png)

The AWARE activity had a planned start date of 15th December 2015, an actual start date of 9th March 2015 and a planned end date of 27th November 2018. Figure 8, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) recording the planned start date.

![Figure 8: Creating the planned start date for DEVCO's AWARE activity](images/devcoPlannedStartDate.png)

Figure 9, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) subsequently retrieving those dates.

![Figure 9: Retrieving the activity dates for DEVCO's AWARE activity](images/devcoReadActivityDates.png)

Figure 10, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) recording AWARE's budget commitment of  US$28,000,000.

![Figure 10: Recording the primary transaction record of DEVCO's AWARE activity](images/devcoTransaction.png)

Figure 11, below, shows [ReportAid](https://github.com/glowkeeper/ReportAid) subsequently retrieving that record.

![Figure 11:  Retrieving the primary transaction of DEVCO's AWARE activity](images/devcoReadTransaction.png)

The AWARE activity recorded on [d-portal](http://d-portal.org) contains more information, such as other budgetary disbursements, administrative contact information and links to other activities related to AWARE. At the time of writing, [ReportAid](https://github.com/glowkeeper/ReportAid) was able to model most of that data, however showing that here would add little to the discussion.

## Analysis

This analysis section constitutes the evaluation and conclusion stage from DSR. The fourth of the subordinate questions of the research objective asks whether blockchains can address criticisms of humanitarian aid? The DSR artefact [ReportAid](https://github.com/glowkeeper/ReportAid), described above through examples, suggests the answer must be yes. Coppi says, "few advanced use cases of blockchain exist in the humanitarian sector. Instead, much discussion relates to potential and anticipated uses of the technology" [@Coppi_BlockchainDistributedLedger_2019]. Although Coppi goes on to describe the use of blockchains in the humanitarian realm, each of the cases he mentions describes direct aid delivery, via the technology's currency capabilities. Therefore, [ReportAid](https://github.com/glowkeeper/ReportAid), as an application that adds trust to transparent aid reporting, is a unique and advanced humanitarian use-case of blockchains that realises some of Coppi's supposed potential of the technology.

Even before the Grand Bargain made at the 2016 World Health Summit (described in the [Literature Review](./literatureReview.md)), the Inter-Agency Standing Committee's Humanitarian Financing Task Team (HFTT) [^707] were busy researching financial transparency because they believed it helped fight corruption through providing the keys to understand, "why, how, what, and how much" [@TransparencyInternational_WhatCorruption_2018]. The HFTT define the '3Ts' of transparent reporting:

1. **Traceability**. The entire transaction chain of aid data must be traceable.
2. **Totality**. Financial information must be complete and relevant.
3. **Timeliness**. Aid information should be up-to-date [@DevelopmentInitiatives_ImprovingHumanitarianTransparency_2017].

However, might the HFTT have missed a 'T', namely, trust? Trust is the glue binding society together because it gives us confidence in situations that might otherwise harbour unknown risks [@Botsman_WhoCanYou_2018]. Indeed, society does not necessarily achieve prosperity as the result of an abundance of natural resources or brilliance of intellect. Neither is that an inevitable result of systemic ideologies, such as frictionless free markets or the communitarian approach of commons-based peer production. Those things may have an important role to play, but ultimately, prosperity, in any form, comes as a result of "spontaneous sociability", achieved through trust because that is the crucial ingredient of any relationship and healthy relationships lead to success [@Fukuyama_TrustSocialVirtues_1996]. [ReportAid](https://github.com/glowkeeper/ReportAid) is an example of an application for humanitarian aid reporting that, by using blockchains, adds that vital ingredient of trust. When an aid organisation adds a record to its implementation of the IATI standard, users can trust that it is that organisation that has created that record because they have digitally signed the transaction that did so.

Furthermore, blockchains address those remaining '3Ts' of transparency, too. After all, its records are publicly viewable and practically impossible to change, and all transactions created are traceable algorithmically [@Savelyev_CopyrightBlockchainEra_2018]. Additionally, the present state of the blockchain is a deterministic function of the genesis block and its ensuing transaction history [@Ethereum_ShardingFAQs_2018]. In other words, a blockchain represent a historical record of all transactions ever recorded on their network, so they are total and timely, too.

## Conclusion

This chapter asks the second of four subordinate questions that help answer the research objective. That asks whether blockchains can address criticisms of humanitarian aid? This chapter examines that fourth question through the lens of the DSR artefact [ReportAid](https://github.com/glowkeeper/ReportAid), which is a blockchain-based application for implementing the IATI standard for aid reporting. The chapter gave some background to [ReportAid](https://github.com/glowkeeper/ReportAid) and described its design. It showed examples of that artefact in use and how it might have been used to document the European Commission's response to the Ebola outbreak in West Africa. That example shows how [ReportAid](https://github.com/glowkeeper/ReportAid) helps address criticism of humanitarian aid by providing the mechanisms for the '4Ts' of transparent humanitarian aid reporting, namely:

1. **Traceability**.
2. **Totality**.
3. **Timeliness**.
4. **Trustworthiness**.

Hence, this chapter concludes that blockchains do indeed contain the necessary properties for addressing criticisms of humanitarian aid.

[^1]: ReportAid is available at the GitHub repository https://github.com/glowkeeper/ReportAid)
[^2]: You can read more about the Agenda for Humanity at https://www.agendaforhumanity.org
[^3]: The IASC Grand Bargain Workstream is available at <https://interagencystandingcommittee.org/greater-transparency>
[^4]: U.N. General Assembly Resolution 46/182 is available at <http://www.un.org/documents/ga/res/46/a46r182.htm>
[^5]: What goes in an organisation file is described at https://iatistandard.org/en/guidance/preparing-data/organisation-information/what-goes-on-your-organisation-file/
[^6]: The activity information that must be published is described at <https://iatistandard.org/en/guidance/preparing-data/activity-information/activity-information-you-can-publish/>
[^7]: d-portal is available at <http://d-portal.org/>
[^211]: More information about the work of HFTT is available at <https://interagencystandingcommittee.org/humanitarian-financing-task-team>
