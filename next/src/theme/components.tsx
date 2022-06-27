import React, { HTMLProps } from 'react'
import { ArrowLeft, ExternalLink as LinkIconFeather, Trash, X } from 'react-feather'
import styled, { keyframes } from 'styled-components'

export const CloseIcon = styled(X)<{ onClick: () => void }>`
  cursor: pointer;
`

export const PageButtons = styled.div`
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 0.2em;
  margin-bottom: 0.5em;
`

export const ArrowIcon = styled.img`
  width: 14px;
  height: 14px;
`

// export const Arrow = styled.div<{ faded: boolean }>`
//   opacity: ${(props) => (props.faded ? 0.3 : 1)};
//   padding: 0 20px;
//   user-select: none;
//   :hover {
//     cursor: pointer;
//   }
// `

export const Break = styled.div`
  height: 1px;
  width: 100%;
`

// for wrapper react feather icons
export const IconWrapper = styled.div<{ stroke?: string; size?: string; marginRight?: string; marginLeft?: string }>`
  display: flex;
  align-items: center;
  justify-content: center;
  width: ${({ size }) => size ?? '20px'};
  height: ${({ size }) => size ?? '20px'};
  margin-right: ${({ marginRight }) => marginRight ?? 0};
  margin-left: ${({ marginLeft }) => marginLeft ?? 0};
  & > * {
    stroke: ${({ theme, stroke }) => stroke ?? theme.white};
  }
`

// A button that triggers some onClick result, but looks like a link.
export const LinkStyledButton = styled.button<{ disabled?: boolean }>`
  border: none;
  text-decoration: none;
  background: none;

  cursor: ${({ disabled }) => (disabled ? 'default' : 'pointer')};
  color: ${({ theme, disabled }) => (disabled ? theme.white : theme.black)};
  font-weight: 500;

  :hover {
    text-decoration: ${({ disabled }) => (disabled ? null : 'underline')};
  }

  :focus {
    outline: none;
    text-decoration: ${({ disabled }) => (disabled ? null : 'underline')};
  }

  :active {
    text-decoration: none;
  }
`

// An internal link from the react-router-dom library that is correctly styled
// export const StyledInternalLink = styled(Link)`
//   text-decoration: none;
//   cursor: pointer;
//   color: ${({ theme }) => theme.primary1};
//   font-weight: 500;
//   background-color: transparent;
//   border: none;
//   shadow: none;

//   :hover {
//     // text-decoration: underline;
//     text-decoration: none;
//   }

//   :focus {
//     outline: none;
//     text-decoration: underline;
//   }

//   :active {
//     text-decoration: none;
//   }
// `

const StyledLink = styled.a`
  text-decoration: none;
  cursor: pointer;
  font-weight: 500;

  :hover {
    text-decoration: underline;
  }

  :focus {
    outline: none;
    text-decoration: underline;
  }

  :active {
    text-decoration: none;
  }
`

const LinkIconWrapper = styled.a`
  text-decoration: none;
  cursor: pointer;
  align-items: center;
  justify-content: center;
  display: flex;

  :hover {
    text-decoration: none;
    opacity: 0.7;
  }

  :focus {
    outline: none;
    text-decoration: none;
  }

  :active {
    text-decoration: none;
  }
`

const LinkIcon = styled(LinkIconFeather)`
  height: 14px;
  width: 18px;
  margin-left: 10px;
  stroke: ${({ theme }) => theme.white};
`

export const TrashIcon = styled(Trash)`
  height: 14px;
  width: 18px;
  margin-left: 10px;
  stroke: ${({ theme }) => theme.black};

  cursor: pointer;
  align-items: center;
  justify-content: center;
  display: flex;

  :hover {
    opacity: 0.7;
  }
`

const rotateImg = keyframes`
  0% {
    transform: perspective(1000px) rotateY(0deg);
  }

  100% {
    transform: perspective(1000px) rotateY(360deg);
  }
`

function navToExternal(event: React.MouseEvent<HTMLAnchorElement>) {
  const { target, href } = event.currentTarget
  if (target === '_blank' || event.ctrlKey || event.metaKey) {
  } else {
    event.preventDefault()
    window.location.href = href
  }
}

/**
 * Outbound link that handles firing google analytics events
 */
export function ExternalLink({
  target = '_blank',
  href,
  rel = 'noopener noreferrer',
  ...rest
}: Omit<HTMLProps<HTMLAnchorElement>, 'as' | 'ref' | 'onClick'> & { href: string }) {
  return <StyledLink target={target} rel={rel} href={href} onClick={navToExternal} {...rest} />
}

export function ExternalLinkIcon({
  target = '_blank',
  href,
  rel = 'noopener noreferrer',
  ...rest
}: Omit<HTMLProps<HTMLAnchorElement>, 'as' | 'ref' | 'onClick'> & { href: string }) {
  return (
    <LinkIconWrapper target={target} rel={rel} href={href} onClick={navToExternal} {...rest}>
      <LinkIcon />
    </LinkIconWrapper>
  )
}

const rotate = keyframes`
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
`

const Spinner = styled.img`
  animation: 4s ${rotate} linear infinite;
  width: 14px;
  height: 14px;
`

export const CustomLightSpinner = styled(Spinner)<{ size: string }>`
  height: ${({ size }) => size};
  width: ${({ size }) => size};
`

export const HideSmall = styled.span`
  ${({ theme }) => theme.mediaWidth.upToSmall`
    display: none;
  `};
`

export const HideExtraSmall = styled.span`
  ${({ theme }) => theme.mediaWidth.upToExtraSmall`
    display: none;
  `};
`

export const SmallOnly = styled.span`
  display: none;
  ${({ theme }) => theme.mediaWidth.upToSmall`
    display: block;
  `};
`
