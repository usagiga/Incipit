/**
 * sleep is to setTimeout as Promise
 * @param msec
 */
export function sleep (msec: number): Promise<void> {
  return new Promise(
    resolve => setTimeout(() => resolve(), msec)
  )
}
