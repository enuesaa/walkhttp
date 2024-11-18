import { Invocation } from '@/gql/types'

type Props = {
  invocation: Invocation
}
export const HistoryDialogHeaders = ({ invocation }: Props) => {
  return (
    <section className='flex py-1'>
      <div className='flex-1'>
        <h4 className='block text-base text-center pb-2 font-medium text-stone-300'>
          Request Headers
        </h4>
        
        <Table headers={invocation.requestHeaders}/>
        
        {invocation.requestHeaders.length === 0 && (
          <div className='text-stone-400 text-center leading-[100px]'>
            The request has no headers.
          </div>
        )}
      </div>

      <div className='flex-1'>
        <h4 className='block text-base text-center pb-2 font-medium text-stone-300'>
          Response Headers
        </h4>

        <Table headers={invocation.responseHeaders}/>

        {invocation.responseHeaders.length === 0 && (
          <div className='text-stone-400 text-center leading-[100px]'>
            The response has no headers.
          </div>
        )}
      </div>
    </section>
  )
}

const Table = ({ headers }: { headers: {name: string; value: string[]}[] }) => {
  return (
    <table className='text-sm w-full'>
      <thead className='text-stone-400'>
        <tr>
          <th className='w-36 font-normal'>Name</th>
          <th className='font-normal'>Value</th>
        </tr>
      </thead>
      <tbody className='bg-stone-800 text-stone-300'>
        {headers.map((h, i) => (
          <tr key={i} className='border-[0.5px] border-l-0 border-r-0 border-stone-700'>
            <td className='py-2 text-center font-semibold'>{h.name}</td>
            <td className='py-2'>{h.value.join(', ')}</td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}
