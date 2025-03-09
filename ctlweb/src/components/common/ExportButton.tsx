import { useSubscribeInvocations } from '@/gql/queries/subscribeInvocations'
import { FaDownload } from 'react-icons/fa'

export const ExportButton = () => {
  const invocations = useSubscribeInvocations()

  const handleClick = () => {  
    if (invocations.fetching || invocations.error !== undefined) {
      return
    }

    const data = {
      invocations: invocations.data?.subscribeInvocations,
    }

    const blob = new Blob([JSON.stringify(data)], {
      type: 'application/json',
    })
    const url = URL.createObjectURL(blob)
    const link = document.createElement('a')
		link.download = 'walkhttp.json'
		link.href = url
		link.click()
		URL.revokeObjectURL(url)
  }

  return (
    <button onClick={handleClick} className='font-normal inline-block h-full px-3 align-middle hover:bg-stone-700'>
      <FaDownload />
    </button>
  )
}
