type Props = {
  label: string
  value: string
}
export const HistoryDialogMetaCard = ({ label, value }: Props) => {
  return (
    <div className='inline-block bg-stone-800 rounded-lg w-fit min-w-24 h-[60px] ml-3 px-2 py-[2px] border border-stone-600'>
      <span className='font-semibold text-stone-400 text-sm'>
        {label}
      </span>
      <div className='w-full text-center text-stone-300 text-base mt-[2px]'>
        {value}
      </div>
    </div>
  )
}
