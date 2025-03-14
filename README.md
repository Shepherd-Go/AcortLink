# Open AcortLink
![GitHub License](https://img.shields.io/github/license/shepherd-go/AcortLink)
![GitHub Release](https://img.shields.io/github/v/release/shepherd-go/AcortLink)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/shepherd-go/AcortLink/integration.yaml)

## Description

Open AcortLink is an open source API where the objective is that any client can connect and shorten any URL, in this API we prioritize stability and speed, using good practices such as testing, 
clean architectures and solid principles, and technologies such as redis, in this release we only have basic functionality but in the following updates we will implement functionality improvements 
and new functions.

## Usage

### Create

POST `https://openacort.link/create`
To create a short URL, only the following parameters are needed.

``` Go
type URLCreate struct {
	URL    string `json:"url" gorm:"column:original_url" validate:"required,url" mold:"trim"`
	Domain string `json:"domain" gorm:"column:domain" validate:"required,url" mold:"trim"`
	Path   string `json:"path" gorm:"column:path" mold:"trim"`
}
```

The 3 parameters expected to create a short URL are:

- `url` (MANDATORY) This field must be in URL format and refer to the URL you want to shorten, example `https://linkedin.com/in/neifer-jesús-reverón-ramos-601451216`


- `domain` (MANDATORY) This is the domain of the client that wants to shorten a URL, that is, the domain from which you are connecting to the API. This is requested only for visual purposes in the
  shortened URL response example `https://tudominio.com/` the / is required at the end of the domain.
  
- `path` (OPTIONAL) This field is for if you want to use a custom path for your shortened URL. If you don't pass this parameter, the API will automatically generate one and send it to you in the response,
  along with the submitted domain.

  ### Example of the request to create a short URL.

``` json
{
    "url":"https://github.com/shepherd-go",
    "domain":"https://tudominio.com/",
    "path":"github"
}
 ```

And the result obtained would be the following:

``` json
{
        "short_url": "https://tudominio.com/github"
}
```

### If the API does not receive a path, it will generate it automatically and an example of this would look like this:

POST `https://openacort.link/create`

``` json
{
    "url":"https://github.com/shepherd-go",
    "domain":"https://tudominio.com/",
}
 ```
Answer obtained. 

 ``` json
{
        "short_url": "https://tudominio.com/x5hb34"
}
```