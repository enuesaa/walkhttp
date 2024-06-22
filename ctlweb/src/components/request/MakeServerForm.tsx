import { Select, TextField, Button, TextArea } from '@radix-ui/themes'
import styles from './MakeForm.css'
import { useMakeServerInvocation } from '@/graph/make-server-invocation'
import { useForm, Controller } from 'react-hook-form'
import { ServerInvocationInput } from '@/graph/types'

export const MakeServerForm = () => {
  const [invoveServerData, invokeServer] = useMakeServerInvocation()

  const {register, handleSubmit, control} = useForm<ServerInvocationInput>()
  const onSubmit = handleSubmit(async(invocation) => await invokeServer({ invocation }))

  return (
    <form className={styles.form} onSubmit={onSubmit}>
      <div className={styles.method}>
        <Controller
          name='method'
          defaultValue='GET'
          control={control}
          render={({field}) => (
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
        <TextField.Root placeholder='https://example.com/' size='3' {...register('url')} />
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
