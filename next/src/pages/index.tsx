import { NFTData } from 'components/NFTCard/components/NFTCard'
import { fetchNFT } from 'data/alchemy'
import { fetchHighestBlock, fetchMints } from 'data/api'
import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import styled from 'styled-components'
import { BlockPageContent } from './block/[number]'

const Home: NextPage<{ mints: NFTData[]; blockNum: string }> = ({ mints, blockNum }) => {
  return <BlockPageContent mints={mints} blockNum={blockNum}></BlockPageContent>
}

export default Home

export async function getServerSideProps() {
  const res1 = await fetchHighestBlock()

  if (res1.error) {
    console.error('res1.error: ', res1.error)
    return {
      props: {
        mints: [],
        blockNum: '',
        error: res1.error,
      },
    }
  }
  if (res1.highestBlock as number) {
    const res = await fetchMints(res1.highestBlock as number)
    const promises = res.mints.map((m) => fetchNFT(m.collection_address, m.token_id))
    const mints = await (await Promise.all(promises)).filter((m) => m != null)
    return { props: { mints, blockNum: `${res1.highestBlock}`, error: null } }
  }
  return {
    props: {
      mints: [],
      blockNum: '',
      error: null,
    },
  }
}
