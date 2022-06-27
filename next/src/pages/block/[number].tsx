import { NFTCard } from 'components/NFTCard'
import { NextPage } from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router'
import styled, { useTheme } from 'styled-components'
import { ArrowIcon } from 'theme'
import ArrowSvg from 'assets/svg/leftArrow.svg'
import Link from 'next/link'
import { NFTData } from 'components/NFTCard/components/NFTCard'
import { fetchMints } from 'data/api'
import { fetchNFT } from 'data/alchemy'

const Background = styled.div`
  position: fixed;
  top: 0;
  left: 0;
  pointer-events: none;
  width: 100vw;
  height: 200vh;
  background: radial-gradient(
    50% 50% at 50% 50%,
    black 0%,
    rgba(252, 187, 7, 0.034) 70%,
    #fcec0804 80%,
    rgba(7, 252, 150, 0.059) 90%,
    rgba(255, 255, 255, 0) 100%
  );
  background: linear-gradient(339deg, #f7690c, #4d4e97);
  // transform: translate(-60vw, -150vh);
  transition: background-image 5s ease-in;
  z-index: 1;
`

const Container = styled.div`
  width: 100vw;
  height: 100%;
  min-height: 100vh;
  // background: linear-gradient(247.34deg, #f7690c -0.44%, #4d4e97 100.3%);
  margin-top: 10rem;
  padding-left: 86px;
  padding-right: 86px;
`

const MintsGrid = styled.div`
  position: relative;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  grid-gap: 2rem;
  gap: 2rem;
  // grid-auto-flow: row dense;
  // max-height: 100;
  overflow-y: scroll;
  overflow-x: visible;
  z-index: 20;
  flex-grow: 1;
  padding-bottom: 2rem;
`

const BlockshotText = styled.div`
  position: fixed;
  top: 0;
  left: 0;
  font-size: 10vw;
  font-weight: bold;
  font-family: 'K2D', sans-serif;
  color: rgba(0, 0, 0, 0.05);
  line-height: 10vw;
  z-index: 2;
`

export function textOutline(width: string, color: string) {
  const widthPx = width
  const positionBase = [`-${widthPx}`, 0, widthPx]
  const positions = positionBase.reduce((arr: any[], x) => arr.concat(positionBase.map((y) => `${x} ${y}`)), [])
  const shadows = positions.map((position: string) => `${position} ${color}`)

  return `text-shadow: ${shadows.join(',')};`
}

const BlockPage: NextPage<{ mints: NFTData[] }> = ({ mints }) => {
  const router = useRouter()
  const block = router.query.number
  return <BlockPageContent mints={mints} blockNum={block as string}></BlockPageContent>
}
export default BlockPage

const RowBetween = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  z-index: 21;
`

const Arrow = styled.a`
  width: 30px;
  rotate: 180deg;
`
const Logo = styled.a`
  stroke: rgba(255, 255, 255, 0.5);
  fill: none;
`

const ArrowRow = styled.div`
  position: fixed;
  top: 30px;
  right: 10px;
  display: flex;
  gap: 5px;
  margin-bottom: 10px;
  z-index: 22;
`
const HeaderRow = styled.div`
  // position: fixed;
  top: 0;
  left: 0;
  display: flex;
  width: 100%;
  max-width: 800px;
  align-items: start;
`

const BlockNumberText = styled.div`
  font-size: 62px;
  line-height: 50px;
  font-weight: bold;
  font-family: 'K2D', sans-serif;
  mix-blend-mode: overlay;
  color: rgba(0, 0, 0, 0.5);
  ${textOutline('2px', 'rgba(255, 255, 255, 0.9)')}
`

export function BlockPageContent({ mints, blockNum }: { mints: NFTData[]; blockNum: string }) {
  return (
    <Container>
      <Head>
        <title>{`Blockshot #${blockNum}`}</title>
      </Head>
      <Background>
        <RowBetween>
          <HeaderRow>
            <Logo href="/">
              <img src={'/logo.svg'} style={{ height: '50px' }} />
            </Logo>
            <BlockNumberText>{`#${blockNum}`}</BlockNumberText>
          </HeaderRow>{' '}
        </RowBetween>
      </Background>
      {blockNum && (
        <ArrowRow>
          <Arrow href={`/block/${parseInt(blockNum) - 1}`}>
            <img src={'/leftArrow.svg'} width={'30px'} />
          </Arrow>
          <Arrow href={`/block/${parseInt(blockNum) + 1}`}>
            <img src={'/rightArrow.svg'} width={'30px'} />
          </Arrow>
        </ArrowRow>
      )}

      <MintsGrid>
        {mints.map((m, i) => {
          return <NFTCard key={`nft-card-${i}`} data={m} />
        })}
      </MintsGrid>
    </Container>
  )
}

export async function getServerSideProps(context: any) {
  const { number } = context.params
  if (number as string) {
    const res = await fetchMints(number)
    const promises = res.mints.map((m) => fetchNFT(m.collection_address, m.token_id))
    const mints = await Promise.all(promises)
    return { props: { mints, error: null } }
  }
  return {
    props: {
      mints: [],
      error: null,
    },
  }
}
