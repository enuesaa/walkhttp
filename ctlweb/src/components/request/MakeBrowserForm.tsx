import { Select, TextField, Button, TextArea } from '@radix-ui/themes'
import { useMakeBrowserInvocation } from '@/gql/queries/makeBrowserInvocation'
import { Controller, useForm } from 'react-hook-form'
import { BrowserInvocationInput } from '@/gql/types'
import { useGetConfig } from '@/gql/queries/getConfig'

export const MakeBrowserForm = () => {
  const appConfig = useGetConfig()
  const [invoveBrowserData, invokeBrowser] = useMakeBrowserInvocation()

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
      <div className='my-2'>
        <Controller
          name='method'
          defaultValue='GET'
          control={control}
          // see https://github.com/orgs/react-hook-form/discussions/8015
          render={({ field: { ref, ...field } }) => (
            <Select.Root size='3' onValueChange={field.onChange} {...field}>
              <Select.Trigger />
              <Select.Content>
                <Select.Item value='GET'>GET</Select.Item>
                <Select.Item value='POST'>POST</Select.Item>
                <Select.Item value='PUT'>PUT</Select.Item>
                <Select.Item value='DELETE'>DELETE</Select.Item>
              </Select.Content>
            </Select.Root>
          )}
        />
      </div>

      <div className='my-2'>
        <TextField.Root defaultValue={appConfig.data?.getConfig.baseUrl} size='3' {...register('url')} />
      </div>

      <div className='my-7 px-2'>
        <span>body</span>
        <TextArea placeholder='{}' {...register('requestBody')}></TextArea>
      </div>

      <div className='text-base mt-2'>
        <Button variant='outline'>Call</Button>
      </div>
    </form>
  )
}
