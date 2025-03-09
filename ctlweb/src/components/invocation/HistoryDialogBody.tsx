import { Invocation } from '@/gql/types'
import { HistoryDialogHeading } from './HistoryDialogHeading'
import { HistoryDialogBodyCopyButton } from './HistoryDialogBodyCopyButton'

type Props = {
  invocation: Invocation
}
export const HistoryDialogBody = ({ invocation }: Props) => {
  return (
    <section className='flex pt-2 mt-5 bg-stone-900'>
      <div className='w-1/2 relative'>
        <HistoryDialogHeading title='Request Body' />
        <Body body={invocation.requestBody} />

        {invocation.requestBody === '' && (
          <div className='text-stone-400 text-center m-12'>The request has no body.</div>
        )}
      </div>

      <div className='w-1/2 relative'>
        <HistoryDialogHeading title='Response Body' />
        <Body body={invocation.responseBody} />

        {invocation.responseBody === '' && (
          <div className='text-stone-400 text-center m-12'>The response has no body.</div>
        )}
      </div>
    </section>
  )
}

const Body = ({ body }: { body: string }) => {
  if (body === '') {
    return <></>
  }

  return (
    <pre className='text-sm p-2 text-stone-400 overflow-x-scroll'>
      <code>{body}</code>
      <HistoryDialogBodyCopyButton text={body} />
    </pre>
  )
}
