import { MouseEventHandler, useState } from 'react'
import { FiCheck, FiCopy } from 'react-icons/fi'

type Props = {
  text: string
}
export const HistoryDialogBodyCopyButton = ({ text }: Props) => {
  const [clicked, setClicked] = useState<boolean>(false)

  const handleCopy: MouseEventHandler<HTMLSpanElement> = async (e) => {
    e.preventDefault()
    await globalThis.navigator.clipboard.writeText(text)

    setClicked(true)
    setTimeout(() => {
      setClicked(false)
    }, 3000)
  }

  return (
    <button onClick={handleCopy} className='absolute top-10 right-3 text-lg'>
      {clicked ? <FiCheck /> : <FiCopy />}
    </button>
  )
}
