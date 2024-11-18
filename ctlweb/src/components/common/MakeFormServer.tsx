import { useMakeServerInvocation } from '@/gql/queries/makeServerInvocation'
import { ServerInvocationInput } from '@/gql/types'
import { useGetConfig } from '@/gql/queries/getConfig'
import { MakeFormHeading } from './MakeFormHeading'
import { FormEventHandler } from 'react'

export const MakeFormServer = () => {
  const appConfig = useGetConfig()
  const [, invokeServer] = useMakeServerInvocation()
  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault()
    const formdata = new FormData(e.currentTarget)
    const invocation = {
      headers: [],
      url: `${appConfig.data?.getConfig.baseUrl}${formdata.get('path')}`,
      method: formdata.get('method'),
      requestBody: formdata.get('requestBody'),
    } as ServerInvocationInput
    await invokeServer({ invocation })
  }

  return (
    <form className='max-w-[700px] mx-auto text-lg leading-relaxed' onSubmit={handleSubmit}>
      <select name='method' className='bg-stone-900 px-2 py-1 border-[0.5px] border-stone-700 outline-none'>
        <option value='GET'>GET</option>
        <option value='POST'>POST</option>
        <option value='PUT'>PUT</option>
        <option value='DELETE'>DELETE</option>
      </select>

      <MakeFormHeading title='url' />
      <div className='text-base'>{appConfig.data?.getConfig.baseUrl}</div>
      <input
        type='text'
        className='block w-full bg-stone-900 border-[0.5px] border-stone-700 text-base text-stone-300 py-1 px-2 outline-none'
        defaultValue='/'
        name='path'
      />

      <MakeFormHeading title='body' />
      <textarea
        className='bg-stone-900 border-[0.5px] border-stone-700 text-base block w-full text-stone-300 py-1 px-2 outline-none'
        name='requestBody'
      />

      <button type='submit' className='mt-5 py-1 px-3 border border-stone-500 rounded-lg hover:bg-stone-700'>Call</button>
    </form>
  )
}
