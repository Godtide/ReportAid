interface AccountProps {
  provider: any
}

export const getAccount = async (props: AccountProps) => {

  const provider = props.provider
  let account = ''
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    account = await signer.getAddress()
  }
  return account
}
