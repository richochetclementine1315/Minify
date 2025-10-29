import React, { useState } from 'react'
import axios from 'axios'
import { toast } from 'react-toastify'
import { FiLink, FiCopy, FiCheck, FiClock, FiZap } from 'react-icons/fi'
import './UrlShortener.css'

const UrlShortener = () => {
  const [url, setUrl] = useState('')
  const [customShort, setCustomShort] = useState('')
  const [expiry, setExpiry] = useState(24)
  const [loading, setLoading] = useState(false)
  const [result, setResult] = useState(null)
  const [copied, setCopied] = useState(false)

  // Handle URL shortening
  const handleShorten = async (e) => {
    e.preventDefault()
    
    // Validation
    if (!url.trim()) {
      toast.error('Please enter a URL')
      return
    }

    // Basic URL validation
    const urlPattern = /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/
    if (!urlPattern.test(url)) {
      toast.error('Please enter a valid URL')
      return
    }

    setLoading(true)
    setResult(null)

    try {
      const response = await axios.post('/api/v1', {
        url: url,
        short: customShort || ' ',
        expiry: expiry
      })

      setResult(response.data)
      toast.success('URL shortened successfully!')
      
      // Reset form
      setUrl('')
      setCustomShort('')
      setExpiry(24)
    } catch (error) {
      const errorMsg = error.response?.data?.error || 'Failed to shorten URL'
      toast.error(errorMsg)
      console.error('Error:', error)
    } finally {
      setLoading(false)
    }
  }

  // Copy to clipboard
  const copyToClipboard = () => {
    navigator.clipboard.writeText(result.short)
    setCopied(true)
    toast.success('Copied to clipboard!')
    
    setTimeout(() => {
      setCopied(false)
    }, 2000)
  }

  // Format time remaining
  const formatTime = (minutes) => {
    if (minutes >= 60) {
      const hours = Math.floor(minutes / 60)
      return `${hours} hour${hours > 1 ? 's' : ''}`
    }
    return `${minutes} minute${minutes > 1 ? 's' : ''}`
  }

  return (
    <div className="url-shortener">
      <div className="shortener-container">
        <div className="shortener-card">
          <div className="card-header">
            <FiLink className="header-icon" />
            <h2>Shorten Your URL</h2>
          </div>

          <form onSubmit={handleShorten} className="shortener-form">
            {/* URL Input */}
            <div className="form-group">
              <label htmlFor="url">Enter your long URL</label>
              <div className="input-wrapper">
                <FiLink className="input-icon" />
                <input
                  type="text"
                  id="url"
                  placeholder="https://example.com/very-long-url"
                  value={url}
                  onChange={(e) => setUrl(e.target.value)}
                  className="form-input"
                />
              </div>
            </div>

            {/* Custom Short URL (Optional) */}
            <div className="form-group">
              <label htmlFor="custom">Custom short URL (optional)</label>
              <div className="input-wrapper">
                <FiZap className="input-icon" />
                <input
                  type="text"
                  id="custom"
                  placeholder="my-custom-link"
                  value={customShort}
                  onChange={(e) => setCustomShort(e.target.value)}
                  className="form-input"
                />
              </div>
              <small className="input-hint">Leave empty for auto-generated short URL</small>
            </div>

            {/* Expiry Time */}
            <div className="form-group">
              <label htmlFor="expiry">
                <FiClock className="label-icon" />
                Expiry Time: {expiry} hours
              </label>
              <input
                type="range"
                id="expiry"
                min="1"
                max="168"
                value={expiry}
                onChange={(e) => setExpiry(parseInt(e.target.value))}
                className="range-input"
              />
              <div className="range-labels">
                <span>1h</span>
                <span>1 week</span>
              </div>
            </div>

            {/* Submit Button */}
            <button
              type="submit"
              disabled={loading}
              className={`submit-btn ${loading ? 'loading' : ''}`}
            >
              {loading ? (
                <>
                  <span className="spinner"></span>
                  Shortening...
                </>
              ) : (
                <>
                  <FiZap />
                  Shorten URL
                </>
              )}
            </button>
          </form>

          {/* Result Section */}
          {result && (
            <div className="result-section">
              <div className="result-card">
                <h3>Your shortened URL is ready!</h3>
                
                <div className="result-url">
                  <div className="url-display">
                    <FiLink className="result-icon" />
                    <span>{result.short}</span>
                  </div>
                  <button
                    onClick={copyToClipboard}
                    className="copy-btn"
                    title="Copy to clipboard"
                  >
                    {copied ? <FiCheck /> : <FiCopy />}
                  </button>
                </div>

                <div className="result-details">
                  <div className="detail-item">
                    <FiClock className="detail-icon" />
                    <span>Expires in: {formatTime(result.rate_limit_reset)}</span>
                  </div>
                  <div className="detail-item">
                    <FiZap className="detail-icon" />
                    <span>Requests remaining: {result.rate_limit}</span>
                  </div>
                </div>

                <div className="original-url">
                  <small>Original URL:</small>
                  <p>{result.url}</p>
                </div>
              </div>
            </div>
          )}
        </div>

        {/* Features Section */}
        <div className="features-section">
          <div className="feature">
            <div className="feature-icon">⚡</div>
            <h3>Lightning Fast</h3>
            <p>Get your shortened URL in milliseconds</p>
          </div>
          <div className="feature">
            <div className="feature-icon">🔒</div>
            <h3>Secure & Private</h3>
            <p>Your data is encrypted and protected</p>
          </div>
          <div className="feature">
            <div className="feature-icon">⏰</div>
            <h3>Custom Expiry</h3>
            <p>Set custom expiration times for your links</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default UrlShortener
