import React from 'react'
import { FiLink } from 'react-icons/fi'
import './Header.css'

const Header = () => {
  return (
    <header className="header">
      <div className="header-content">
        <div className="logo">
          <FiLink className="logo-icon" />
          <h1>Minify</h1>
        </div>
        <p className="tagline">Shorten your URLs, amplify your reach</p>
      </div>
    </header>
  )
}

export default Header
