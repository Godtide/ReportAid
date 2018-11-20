interface AccountProps {
  provider: any
}

export const getAccount = (props: AccountProps) => {

  console.log('In set account')
  let providerAccount = ''
  const provider = props.provider
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    signer.getAddress().then((account: string) => {
        providerAccount = account
        console.log ('Setting account', account)
    }).catch((error: any) => {
      console.log(error)
    })
  }

  return providerAccount
}
