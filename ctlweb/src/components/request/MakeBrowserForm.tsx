import { Select, TextField, Button, TextArea } from '@radix-ui/themes'
import styles from './MakeForm.css'
import { useMakeBrowserInvocation } from '@/gql/queries/makeBrowserInvocation'
import { Controller, useForm } from 'react-hook-form'
import { BrowserInvocationInput } from '@/graph/types'
import { useGetAppConfig } from '@/gql/queries/appConfig'

export const MakeBrowserForm = () => {
  const appConfig = useGetAppConfig()
  const [invoveBrowserData, invokeBrowser] = useMakeBrowserInvocation()

  const { register, handleSubmit, control } = useForm<BrowserInvocationInput>()
  const onSubmit = handleSubmit(async (invocation) => {
    invocation.requestHeaders = []
    invocation.status = 0
    invocation.responseHeaders = []
    invocation.responseBody = ''

    try {
      const res = await fetch(invocation.url, {
        method: invocation.method,
        body: invocation.requestBody,
      })
      invocation.status = res.status
    } catch (e) {
      console.log(e)
    }

    await invokeBrowser({ invocation })
  })

  return (
    <form className={styles.form} onSubmit={onSubmit}>
      <div className={styles.method}>
        <Controller
          name='method'
          defaultValue='GET'
          control={control}
          render={({ field }) => (
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

      <div className={styles.url}>
        <TextField.Root defaultValue={appConfig.data?.appConfig.baseUrl} size='3' {...register('url')} />
      </div>

      <div className={styles.body}>
        <span>body</span>
        <TextArea placeholder='{}' {...register('requestBody')}></TextArea>
      </div>

      <div className={styles.btn}>
        <Button>Call</Button>
      </div>
    </form>
  )
}
