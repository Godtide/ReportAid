import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getDocs } from '../../../store/IATI/IATIReader/organisations/organisationDocs/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationDocsData } from '../../../store/IATI/IATIReader/types'

import { OrganisationDoc as OrganisationDocStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrganisationDocProps {
  num: number
  orgDocs: OrganisationDocsData
}

interface OrganisationDocDispatchProps {
  getDocs: () => void
}

type OrganisationDocsReaderProps =  WithStyles<typeof styles> & OrganisationDocProps & OrganisationDocDispatchProps

class Docs extends React.Component<OrganisationDocsReaderProps> {

  constructor (props: OrganisationDocsReaderProps) {
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
        xs += `**${OrganisationDocStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgDocs[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrganisationDocStrings.numDocs}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((docKey) => {
          //console.log('Doc: ', values[1][docKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][docKey].hasOwnProperty('category') && values[1][docKey].category != "" ) {
            const category = ethers.utils.parseBytes32String(values[1][docKey].category)
            const countryRef = ethers.utils.parseBytes32String(values[1][docKey].countryRef)
            const lang = ethers.utils.parseBytes32String(values[1][docKey].lang)
            const date = ethers.utils.parseBytes32String(values[1][docKey].date)
            xs+= `**${OrganisationDocStrings.reportingOrgRef}**: ${values[1][docKey].report.orgRef} <br />`
            xs+= `**${OrganisationDocStrings.docReference}**: ${docKey} <br />`
            xs+= `**${OrganisationDocStrings.documentTitle}**: ${values[1][docKey].title} <br />`
            xs+= `**${OrganisationDocStrings.documentFormat}**: ${values[1][docKey].format} <br />`
            xs+= `**${OrganisationDocStrings.documentURL}**: ${values[1][docKey].url} <br />`
            xs+= `**${OrganisationDocStrings.documentCategory}**: ${category} <br />`
            xs+= `**${OrganisationDocStrings.documentCountryRef}**: ${countryRef} <br />`
            xs+= `**${OrganisationDocStrings.documentDesc}**: ${values[1][docKey].desc} <br />`
            xs+= `**${OrganisationDocStrings.documentLang}**: ${lang} <br />`
            xs+= `**${OrganisationDocStrings.documentDate}**: ${date} <br /><br />`
          }
        })
        length += 1
        length == docsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrganisationDocStrings.headingOrganisationDocReader}</h2>
        <p>
          <b>{OrganisationDocStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrganisationDocStrings.reportDocDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationDocProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgDocsReader.num,
    orgDocs: state.orgDocsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationDocDispatchProps => {
  return {
    getDocs: () => dispatch(getDocs())
  }
}

export const OrganisationDocs = withTheme(withStyles(styles)(connect<OrganisationDocProps, OrganisationDocDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Docs)))
