import { Reducer } from 'redux'
import { InfoPageProps } from './types'
import { About } from '../../utils/strings'
import { Help } from '../../utils/strings'
import { Home } from '../../utils/strings'
import { IATIReader } from '../../utils/strings'
import { IATIWriter } from '../../utils/strings'
import { Overview } from '../../utils/strings'

const initialState: InfoPageProps = {
  data: {
    about: {
      title: About.heading,
      data: About.info
    },
    help: {
      title: Help.heading,
      data: Help.info
    },
    home: {
      title: Home.heading,
      data: Home.info
    },
    IATIReader: {
      title: IATIReader.heading,
      data: IATIReader.info
    },
    IATIWriter: {
      title: IATIWriter.heading,
      data: IATIWriter.info
    },
    overview: {
      title: Overview.heading,
      data: Overview.info
    }
  }
}

export const reducer: Reducer<InfoPageProps> = (state = initialState): InfoPageProps => {
  return state
}
