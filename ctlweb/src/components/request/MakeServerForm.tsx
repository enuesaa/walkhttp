import { Select, TextField, Button, TextArea } from '@radix-ui/themes'
import styles from './MakeForm.css'
import { useMakeServerInvocation } from '@/graph/make-server-invocation'
import { FormEventHandler } from 'react'
import { useForm } from '@tanstack/react-form'
import type { FieldApi } from '@tanstack/react-form'

export const MakeServerForm = () => {
  const [invoveServerData, invokeServer] = useMakeServerInvocation()

  const form = useForm({
    defaultValues: {
      url: '',
    },
    onSubmit: async ({ value }) => {
      console.log(value)
      // const invocation = {
      //   method: (e.currentTarget.elements.namedItem('method') as HTMLInputElement).value,
      //   url: (e.currentTarget.elements.namedItem('url') as HTMLInputElement).value,
      //   requestHeaders: [],
      //   requestBody: (e.currentTarget.elements.namedItem('body') as HTMLInputElement).value,
      // }
      // await invokeServer({ invocation })
    },
  })

  return (
    <form className={styles.form} onSubmit={e => {e.preventDefault();form.handleSubmit()}}>
      <div className={styles.method}>
        <Select.Root defaultValue='GET' size='3' name='method'>
          <Select.Trigger />
          <Select.Content>
            <Select.Item value='GET'>GET</Select.Item>
            <Select.Item value='POST'>POST</Select.Item>
            <Select.Item value='PUT'>PUT</Select.Item>
            <Select.Item value='DELETE'>DELETE</Select.Item>
          </Select.Content>
        </Select.Root>
      </div>

      <div className={styles.url}>
        <form.Field
          name='url'
          children={f => (
            <TextField.Root name={f.name} placeholder='https://example.com/' size='3' onChange={(e) => f.handleChange(e.target.value)} />
          )}
        />
      </div>
{/* 
      <div className={styles.body}>
        <span>body</span>
        <TextArea name='body' placeholder='{}'></TextArea>
      </div> */}

      <div className={styles.btn}>
        <form.Subscribe
          children={(<Button type='submit'>Call</Button>)}
        />
      </div>
    </form>
  )
}
