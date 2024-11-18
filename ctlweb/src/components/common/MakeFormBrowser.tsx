import { useMakeBrowserInvocation } from '@/gql/queries/makeBrowserInvocation'
import { useForm } from 'react-hook-form'
import { BrowserInvocationInput } from '@/gql/types'
import { useGetConfig } from '@/gql/queries/getConfig'
import { MakeFormHeading } from './MakeFormHeading'

export const MakeFormBrowser = () => {
  const appConfig = useGetConfig()
  const [, invokeBrowser] = useMakeBrowserInvocation()

  const { register, handleSubmit, control } = useForm<BrowserInvocationInput>()
  const onSubmit = handleSubmit(async (invocation) => {
    invocation.requestHeaders = []
    invocation.status = 0
    invocation.responseHeaders = []
    invocation.responseBody = ''
    invocation.started = new Date().toISOString()

    try {
      const res = await fetch(invocation.url, {
        method: invocation.method,
        body: invocation.method === 'GET' ? undefined : invocation.requestBody,
      })
      invocation.received = new Date().toISOString()
      invocation.status = res.status
      for (const [name, value] of res.headers.entries()) {
        invocation.responseHeaders.push({ name, value: [value] })
      }
    } catch (e) {
      console.log(e)
    }

    await invokeBrowser({ invocation })
  })

  return (
    <form className='max-w-[700px] mx-auto text-lg leading-relaxed' onSubmit={onSubmit}>
      <select {...register('method')} className='bg-stone-900 px-2 py-1 border-[0.5px] border-stone-700 outline-none'>
        <option value='GET'>GET</option>
        <option value='POST'>POST</option>
        <option value='PUT'>PUT</option>
        <option value='DELETE'>DELETE</option>
      </select>

      <MakeFormHeading title='url' />
      <input
        type='text'
        className='bg-stone-900 border-[0.5px] border-stone-700 text-base block w-full text-stone-300 py-1 px-2 outline-none'
        defaultValue={appConfig.data?.getConfig.baseUrl}
        {...register('url')}
      />

      <MakeFormHeading title='body' />
      <textarea
        className='bg-stone-900 border-[0.5px] border-stone-700 text-base block w-full text-stone-300 py-1 px-2 outline-none'
        {...register('requestBody')}
      />

      <button type='submit' className='mt-5 py-1 px-3 border border-stone-500 rounded-lg hover:bg-stone-700'>Call</button>
    </form>
  )
}
