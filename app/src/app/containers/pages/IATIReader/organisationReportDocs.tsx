import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getReportDocs } from '../../../store/IATI/IATIReader/organisationReportDocs/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgReportDocsData } from '../../../store/IATI/IATIReader/organisationReportDocs/types'

import { OrganisationReportDoc as OrgReportDocStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgReportDocProps {
  num: number
  orgReportDocs: OrgReportDocsData
}

interface OrgReportDocDispatchProps {
  getReportDocs: () => void
}

type OrgReportDocsReaderProps =  WithStyles<typeof styles> & OrgReportDocProps & OrgReportDocDispatchProps

export class OrgReportDocs extends React.Component<OrgReportDocsReaderProps> {

  constructor (props: OrgReportDocsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getReportDocs()
  }

  render() {

    const docsData = Object.keys(this.props.orgReportDocs)
    let xs = ""
    if ( docsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      docsData.forEach((reportKey) => {
        xs += `**${OrgReportDocStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgReportDocs[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgReportDocStrings.numReportDocs}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((docKey) => {
          //console.log('Doc: ', values[1][docKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][docKey].title != "" ) {
            const category = ethers.utils.parseBytes32String(values[1][docKey].category)
            const countryRef = ethers.utils.parseBytes32String(values[1][docKey].countryRef)
            const lang = ethers.utils.parseBytes32String(values[1][docKey].lang)
            const date = ethers.utils.parseBytes32String(values[1][docKey].date)
            xs+= `**${OrgReportDocStrings.reportingOrgRef}**: ${values[1][docKey].report.orgRef} <br />`
            xs+= `**${OrgReportDocStrings.docReference}**: ${docKey} <br />`
            xs+= `**${OrgReportDocStrings.documentTitle}**: ${values[1][docKey].title} <br />`
            xs+= `**${OrgReportDocStrings.documentFormat}**: ${values[1][docKey].format} <br />`
            xs+= `**${OrgReportDocStrings.documentURL}**: ${values[1][docKey].url} <br />`
            xs+= `**${OrgReportDocStrings.documentCategory}**: ${category} <br />`
            xs+= `**${OrgReportDocStrings.documentCountryRef}**: ${countryRef} <br />`
            xs+= `**${OrgReportDocStrings.documentDesc}**: ${values[1][docKey].desc} <br />`
            xs+= `**${OrgReportDocStrings.documentLang}**: ${lang} <br />`
            xs+= `**${OrgReportDocStrings.documentDate}**: ${date} <br /><br />`
          }
        })
        length += 1
        length == docsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgReportDocStrings.headingOrgReportDocReader}</h2>
        <p>
          <b>{OrgReportDocStrings.numReports}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgReportDocStrings.reportDocDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportDocProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReportDocsReader.num,
    orgReportDocs: state.orgReportDocsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportDocDispatchProps => {
  return {
    getReportDocs: () => dispatch(getReportDocs())
  }
}

export const OrganisationReportDocs = withTheme(withStyles(styles)(connect<OrgReportDocProps, OrgReportDocDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportDocs)))
