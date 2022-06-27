import { configureStore } from '@reduxjs/toolkit'
import { setupListeners } from '@reduxjs/toolkit/query/react'
import { load, save } from 'redux-localstorage-simple'

import application from './application/reducer'

import user from './user/reducer'

const store = configureStore({
  reducer: {
    application,
    user,
    // [dataApi.reducerPath]: dataApi.reducer,
  },
  middleware: (getDefaultMiddleware: any) => getDefaultMiddleware({ thunk: true, serializableCheck: false }),
  // .concat(dataApi.middleware),
})

setupListeners(store.dispatch)

export default store

export type AppState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
