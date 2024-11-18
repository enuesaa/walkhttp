import { useMakeServerInvocation } from '@/gql/queries/makeServerInvocation'
import { useForm } from 'react-hook-form'
import { ServerInvocationInput } from '@/gql/types'
import { useGetConfig } from '@/gql/queries/getConfig'
import { MakeFormHeading } from './MakeFormHeading'

export const MakeFormServer = () => {
  const appConfig = useGetConfig()
  const [, invokeServer] = useMakeServerInvocation()

  const { register, handleSubmit } = useForm<ServerInvocationInput>()
  const onSubmit = handleSubmit(async (invocation) => await invokeServer({ invocation }))

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
