import { useCallback } from 'react'
import { useAppDispatch, useAppSelector } from 'state/hooks'

import { AppState } from '../index'
import { ApplicationModal, setOpenModal } from './reducer'

export function useModalOpen(modal: ApplicationModal | string): boolean {
  const openModal = useAppSelector((state: AppState) => state.application.openModal)
  return openModal === modal
}

export function useToggleModal(modal: ApplicationModal | string): () => void {
  const open = useModalOpen(modal)
  console.log('mouse useToggleModal. open:', open)
  const dispatch = useAppDispatch()
  return useCallback(() => dispatch(setOpenModal(open ? null : modal)), [dispatch, modal, open])
}
