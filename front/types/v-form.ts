/**
 * VForm is interface of form in $refs
 */
export interface VForm {
  validate(): boolean

  reset(): void

  resetValidation(): void
}
