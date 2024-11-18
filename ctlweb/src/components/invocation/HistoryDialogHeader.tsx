import { Invocation, Header } from '@/gql/types'
import { HistoryDialogHeading } from './HistoryDialogHeading'

type Props = {
  invocation: Invocation
}
export const HistoryDialogHeader = ({ invocation }: Props) => {
  return (
    <section className='flex pt-2 gap-1'>
      <div className='w-1/2'>
        <HistoryDialogHeading title='Request Headers' />
        <Table headers={invocation.requestHeaders} />

        {invocation.requestHeaders.length === 0 && (
          <div className='text-stone-400 text-center mt-16'>
            The request has no headers.
          </div>
        )}
      </div>

      <div className='w-1/2'>
        <HistoryDialogHeading title='Response Headers' />
        <Table headers={invocation.responseHeaders} />

        {invocation.responseHeaders.length === 0 && (
          <div className='text-stone-400 text-center mt-16'>
            The response has no headers.
          </div>
        )}
      </div>
    </section>
  )
}

const Table = ({ headers }: { headers: Header[] }) => {
  return (
    <table className='text-sm w-full text-stone-400'>
      <tbody>
        {headers.map((h, i) => (
          <tr key={i} className='border-[0.5px] border-l-0 border-r-0 border-stone-700'>
            <td className='py-2 text-center font-semibold w-44'>{h.name}</td>
            <td className='py-2 text-stone-300 break-all px-2'>{h.value.join(', ')}</td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}
