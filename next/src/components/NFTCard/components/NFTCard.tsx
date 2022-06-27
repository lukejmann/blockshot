import { useState, useEffect, CSSProperties } from 'react'
import styled, { useTheme } from 'styled-components'
import Media from './Media'
import { T } from 'theme'

const NFTCardContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  width: fit-content;
`

export type NFTData = {
  tokenId: string
  media: {
    gateway: string
  }[]
  metadata: {
    name: string
    image: string
  }
}
const RowBetween = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
`

const Column = styled.div`
  display: flex;
  flex-direction: column;
`
export function NFTCard({ data }: { data: NFTData }) {
  const theme = useTheme()

  if (!data || data?.media[0].gateway == '') return null

  return (
    <NFTCardContainer>
      <Column>
        <Media media={data?.media[0].gateway ?? ''} autoPlay></Media>
        <RowBetween>
          <div></div>
          <T.main color={theme.white}>{data?.metadata.name}</T.main>
        </RowBetween>
      </Column>
    </NFTCardContainer>
  )
}
