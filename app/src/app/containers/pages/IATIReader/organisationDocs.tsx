import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getDocs } from '../../../store/IATI/IATIReader/organisations/organisationDocs/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgDocsData } from '../../../store/IATI/IATIReader/organisationDocs/types'

import { OrganisationDoc as OrgDocStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgDocProps {
  num: number
  orgDocs: OrgDocsData
}

interface OrgDocDispatchProps {
  getDocs: () => void
}

type OrgDocsReaderProps =  WithStyles<typeof styles> & OrgDocProps & OrgDocDispatchProps

export class OrgDocs extends React.Component<OrgDocsReaderProps> {

  constructor (props: OrgDocsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getDocs()
  }

  render() {

    const docsData = Object.keys(this.props.orgDocs)
    let xs = ""
    if ( docsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      docsData.forEach((reportKey) => {
        xs += `**${OrgDocStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgDocs[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgDocStrings.numDocs}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((docKey) => {
          //console.log('Doc: ', values[1][docKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][docKey].hasOwnProperty('category') && values[1][docKey].category != "" ) {
            const category = ethers.utils.parseBytes32String(values[1][docKey].category)
            const countryRef = ethers.utils.parseBytes32String(values[1][docKey].countryRef)
            const lang = ethers.utils.parseBytes32String(values[1][docKey].lang)
            const date = ethers.utils.parseBytes32String(values[1][docKey].date)
            xs+= `**${OrgDocStrings.reportingOrgRef}**: ${values[1][docKey].report.orgRef} <br />`
            xs+= `**${OrgDocStrings.docReference}**: ${docKey} <br />`
            xs+= `**${OrgDocStrings.documentTitle}**: ${values[1][docKey].title} <br />`
            xs+= `**${OrgDocStrings.documentFormat}**: ${values[1][docKey].format} <br />`
            xs+= `**${OrgDocStrings.documentURL}**: ${values[1][docKey].url} <br />`
            xs+= `**${OrgDocStrings.documentCategory}**: ${category} <br />`
            xs+= `**${OrgDocStrings.documentCountryRef}**: ${countryRef} <br />`
            xs+= `**${OrgDocStrings.documentDesc}**: ${values[1][docKey].desc} <br />`
            xs+= `**${OrgDocStrings.documentLang}**: ${lang} <br />`
            xs+= `**${OrgDocStrings.documentDate}**: ${date} <br /><br />`
          }
        })
        length += 1
        length == docsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgDocStrings.headingOrgDocReader}</h2>
        <p>
          <b>{OrgDocStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgDocStrings.reportDocDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgDocProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgDocsReader.num,
    orgDocs: state.orgDocsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDocDispatchProps => {
  return {
    getDocs: () => dispatch(getDocs())
  }
}

export const OrganisationDocs = withTheme(withStyles(styles)(connect<OrgDocProps, OrgDocDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgDocs)))
