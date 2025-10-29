import React from 'react'
import { FiGithub, FiHeart } from 'react-icons/fi'
import './Footer.css'

const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-content">
        <p>
          Made with <FiHeart className="heart-icon" /> by Mrinmoy Matilal
        </p>
        <div className="footer-links">
          <a 
            href="https://github.com/richochetclementine1315/Minify" 
            target="_blank" 
            rel="noopener noreferrer"
            className="github-link"
          >
            <FiGithub /> GitHub
          </a>
        </div>
      </div>
    </footer>
  )
}

export default Footer
