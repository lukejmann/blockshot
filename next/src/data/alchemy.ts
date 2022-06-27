import { ApolloClient, NormalizedCacheObject } from '@apollo/client'
import { NFTData } from 'components/NFTCard/components/NFTCard'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import weekOfYear from 'dayjs/plugin/weekOfYear'
import gql from 'graphql-tag'
import styled, { useTheme } from 'styled-components'

export const fetchNFT = async (collectionAddress: string, tokenId: string) => {
  const apiKey = process.env.NEXT_PUBLIC_ALCHEMY_API_KEY
  const baseURL = `https://eth-mainnet.alchemyapi.io/nft/v2/${apiKey}/getNFTMetadata`
  const fetchURL = `${baseURL}?contractAddress=${collectionAddress}&tokenId=${tokenId}`
  const res = await fetch(fetchURL, {
    method: 'GET',
    redirect: 'follow',
  })
  try {
    const data = await res.json()
    return data as NFTData
  } catch (e) {
    console.log(e)
    return null
  }
}
