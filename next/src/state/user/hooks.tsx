import { useCallback } from 'react'
import { shallowEqual } from 'react-redux'
import { useAppDispatch, useAppSelector } from 'state/hooks'

import { updateUserDarkMode } from './actions'

export function useIsDarkMode(): boolean {
  const { userDarkMode } = useAppSelector(
    ({ user: { userDarkMode } }) => ({
      userDarkMode,
    }),
    shallowEqual
  )

  return userDarkMode ?? false
}

export function useDarkModeManager(): [boolean, () => void] {
  const dispatch = useAppDispatch()
  const darkMode = useIsDarkMode()

  const toggleSetDarkMode = useCallback(() => {
    dispatch(updateUserDarkMode({ userDarkMode: !darkMode }))
  }, [darkMode, dispatch])

  return [darkMode, toggleSetDarkMode]
}
