import React, { useEffect, useState } from 'react'
import styled, { CSSProperties } from 'styled-components'
import Loading from './Loading'

const mediaContentStyle: CSSProperties = {
  width: '100%',
  height: '100%',
  objectFit: 'contain',
}

function Text({ media }: { media: string }) {
  const [content, setContent] = useState<string | null>(null)

  useEffect(() => {
    fetch(media)
      .then((r) => r.text())
      .then((r) => setContent(r))
  }, [])

  return <div style={mediaContentStyle}>{content}</div>
}

function Video({ media, autoPlay }: { media: string; autoPlay: boolean }) {
  return (
    <video style={mediaContentStyle} muted autoPlay={autoPlay} controls={!autoPlay} loop playsInline>
      <source src={media} />
    </video>
  )
}

function Audio({ media }: { media: string }) {
  return <audio style={mediaContentStyle} controls src={media}></audio>
}

const MediaContainer = styled.div`
  max-width: 400px;
  max-height: 400px;
  overflow: hidden;
  border-radius: 5px;
  box-shadow: -4.10233px 5.46977px 31.4512px 5.46977px rgba(0, 0, 0, 0.25);
`

export default function Media({ media, autoPlay }: { media: string; autoPlay: boolean }) {
  const content = () => {
    if (media?.includes('mp4')) return <Video media={media} autoPlay={autoPlay} />

    if (media?.includes('audio')) return <Audio media={media} />
    return <img src={media} style={mediaContentStyle} />
  }

  return !media ? <Loading></Loading> : <MediaContainer>{content()}</MediaContainer>
}
