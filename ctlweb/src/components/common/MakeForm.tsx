import { SegmentedControl } from '@radix-ui/themes'
import { useState } from 'react'
import { MakeFormBrowser } from './MakeFormBrowser'
import { MakeFormServer } from './MakeFormServer'

export const MakeForm = () => {
  const [fromServer, setFromServer] = useState(true)

  const handleChangeFrom = (from: string) => {
    setFromServer(from === 'Server')
  }

  return (
    <>
      <div className='text-center'>
        <SegmentedControl.Root defaultValue='Server' size='3' radius='full' onValueChange={handleChangeFrom}>
          <SegmentedControl.Item value='Server'>Server</SegmentedControl.Item>
          <SegmentedControl.Item value='Browser'>Browser</SegmentedControl.Item>
        </SegmentedControl.Root>
      </div>
      {fromServer ? <MakeFormServer /> : <MakeFormBrowser />}
    </>
  )
}
