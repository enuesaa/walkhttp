export const fmtdate = (datestr: string): string => {
  const datefmt = new Intl.DateTimeFormat('en-US', {
    hour: 'numeric',
    hour12: false,
    minute: 'numeric',
    second: 'numeric',
  })
  try {
    const date = new Date(datestr)
    return datefmt.format(date)
  } catch {
    return datestr
  }
}

export const calcDuration = (old: string, future: string): number => {
  try {
    const olddt = new Date(old)
    const futuredt = new Date(future)
    return futuredt.getTime() - olddt.getTime()
  } catch {
    return 0
  }
}
