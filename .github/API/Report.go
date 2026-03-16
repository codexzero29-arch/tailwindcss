{
  "title": "GitHub Pages",
  "description": "The configuration for GitHub Pages for a repository.",
  "type": "object",
  "properties": {
    "url": {
      "type": "string",
      "description": "The API address for accessing this Page resource.",
      "format": "uri",
      "examples": [
        "https://api.github.com/repos/github/hello-world/pages"
      ]
    },
    "status": {
      "type": [
        "string",
        "null"
      ],
      "description": "The status of the most recent build of the Page.",
      "enum": [
        "built",
        "building",
        "errored",
        null
      ],
      "examples": [
        "built"
      ]
    },
    "cname": {
      "description": "The Pages site's custom domain",
      "type": [
        "string",
        "null"
      ],
      "examples": [
        "example.com"
      ]
    },
    "protected_domain_state": {
      "type": [
        "string",
        "null"
      ],
      "description": "The state if the domain is verified",
      "enum": [
        "pending",
        "verified",
        "unverified",
        null
      ],
      "examples": [
        "pending"
      ]
    },
    "pending_domain_unverified_at": {
      "type": [
        "string",
        "null"
      ],
      "description": "The timestamp when a pending domain becomes unverified.",
      "format": "date-time"
    },
    "custom_404": {
      "type": "boolean",
      "description": "Whether the Page has a custom 404 page.",
      "default": false,
      "examples": [
        false
      ]
    },
    "html_url": {
      "type": "string",
      "description": "The web address the Page can be accessed from.",
      "format": "uri",
      "examples": [
        "https://example.com"
      ]
    },
    "build_type": {
      "type": [
        "string",
        "null"
      ],
         "description": "The process in which the Page will be built.",
      "enum": [
        "legacy",
        "workflow",
        null
      ],
      "examples": [
        "legacy"
      ]
    },
    "source": {
      "title": "Pages Source Hash",
      "type": "object",
      "properties": {
        "branch": {
          "type": "string"
        },
        "path": {
          "type": "string"
        }
      },
      "required": [
        "branch",
        "path"
      ]
    },
    "public": {
      "type": "boolean",
      "description": "Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.",
      "examples": [
        true
      ]
    },
    "https_certificate": {
      "title": "Pages Https Certificate",
      "type": "object",
      "properties": {
        "state": {
          "type": "string",
          "enum": [
            "new",
            "authorization_created",
            "authorization_pending",
            "authorized",
            "authorization_revoked",
            "issued",
            "uploaded",
            "approved",
            "errored",
            "bad_authz",
            "destroy_pending",
            "dns_changed"
          ],
          "examples": [
            "approved"
          ]
        },
        "description": {
          "type": "string",
          "examples": [
            "Certificate is approved"
          ]
        },
        "domains": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Array of the domain set and its alternate name (if it is configured)",
          "examples": [
            "example.com",
            "www.example.com"
          ]
        },
        "expires_at": {
          "type": "string",
          "format": "date"
        }
      },
      "required": [
        "state",
        "description",
        "domains"
      ]
    },
    "https_enforced": {
      "type": "boolean",
      "description": "Whether https is enabled on the domain",
      "examples": [
        true
      ]
    }
  },
  "required": [
    "url",
    "status",
    "cname",
    "custom_404",
    "public"
   ]
  }

