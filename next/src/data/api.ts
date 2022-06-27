import { ApolloClient, NormalizedCacheObject } from '@apollo/client'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import weekOfYear from 'dayjs/plugin/weekOfYear'
import gql from 'graphql-tag'
import styled, { useTheme } from 'styled-components'

// format dayjs with the libraries that we need
dayjs.extend(utc)
dayjs.extend(weekOfYear)

const AutoRow = styled.div`
  display: flex;
  width: 100%;
`

export interface Mint {
  collection_address: string
  token_id: string
}

export const PAGE_SIZE = 10

export async function fetchMints(blockNumber: number): Promise<{
  mints: Mint[]
  error: string | null
}> {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/block/` + blockNumber, {
      method: 'GET',
      headers: {
        'X-Requested-With': 'XMLHttpRequest',
      },
    })
    const data = await res.json()

    return {
      mints: data as Mint[],
      error: null,
    }
  } catch (e) {
    console.log(e)
    return {
      mints: [],
      error: e.message,
    }
  }
}

export async function fetchHighestBlock(): Promise<{
  highestBlock: number | null
  error: string | null
}> {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/highest`, {
      method: 'GET',
      headers: {
        'X-Requested-With': 'XMLHttpRequest',
      },
    })
    const data = await res.json()
    return {
      highestBlock: data as number,
      error: null,
    }
  } catch (e) {
    console.log(e)
    return {
      highestBlock: null,
      error: e.message,
    }
  }
}
