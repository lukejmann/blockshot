import { createReducer } from '@reduxjs/toolkit'

import { updateUserDarkMode } from './actions'

const currentTimestamp = () => new Date().getTime()

export interface UserState {
  userDarkMode: boolean | null
}

export const initialState: UserState = {
  userDarkMode: null,
}

export default createReducer(initialState, (builder) =>
  builder.addCase(updateUserDarkMode, (state, action) => {
    state.userDarkMode = action.payload.userDarkMode
  })
)
