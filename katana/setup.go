package katana

import (
  "strconv"

  "github.com/nasciiboy/txt"
)

func (d *doc) SetupHunter() {
  for d.Rune != EOF {
    switch whoIsThere( d.Line ) {
    case NodeBlock:
      if block := d.GetBlock(); block != nil {
        d.setSetupCommand( block )
      } else { d.NextLine() }
    case NodeComment: d.NextLine()
    case NodeEmpty: return
    default: return
    }
  }

  return
}

func (d *doc) setSetupCommand( block *Block ){
  switch block.Comm.Text {
  case "title"      : d.Title          = block.Head.GetFancyMarkup()
  case "subtitle"   : d.Subtitle       = block.Head.GetFancyMarkup()
  case "author"     : d.Author         = append( d.Author, txt.RmSpacesToTheSides( block.Head.Text() ))
  case "style"      : d.Style          = append( d.Style,  txt.RmSpacesToTheSides( block.Head.Text() ))
  case "translator" : d.Translator     = append( d.Translator, txt.RmSpacesToTheSides( block.Head.Text() ))
  case "source"     : d.Source         = append( d.Source, txt.RmSpacesToTheSides( block.Head.Text() ))
  case "licence"    : d.Licence        = txt.RmSpacesToTheSides( block.Head.Text() )
  case "id"         : d.ID             = txt.RmSpacesToTheSides( block.Head.Text() )
  case "date"       : d.Date           = txt.RmSpacesToTheSides( block.Head.Text() )
  case "tags"       : d.Tags           = append( d.Tags, block.Head.getTags()... )
  case "options"    : d.setSetupOptions( block.Head.GetArgs() )
  case "description": d.Description    = txt.Linelize( block.Head.Text() )
  case "lang"       : d.Lang           = txt.RmSpacesToTheSides( block.Head.Text() )
  case "language"   : d.Lang           = txt.RmSpacesToTheSides( block.Head.Text() )
  default:
    d.MoreConfigs[ block.Comm.Text ] = append( d.MoreConfigs[ block.Comm.Text ], block.Head.Text() )
  }
}

var defaultOpts = map[string][]Arg {
  "mathJax"  : {{ "", String }},
  "fancyCode": {{ "", String }},
  "hShift"   : {{ "0", Int }},
  "toc"      : {{ "true", Bool }},
}

func (d *doc) setSetupOptions( args []ArgType ){
  for _, opt := range args {
    copt, cannon := d.syncOpt( opt, defaultOpts )
    if !cannon {
      d.MultOptions[ opt.Name ] = opt.Args
      continue
    }

    switch copt.Name {
    case "fancyCode":
      d.BoolOptions["fancyCode"] = true
      d.TextOptions["fancyCode"] = copt.Args[0].Data
    case "hShif":
      i64, _ := strconv.ParseInt( copt.Args[0].Data, 10, 8 )
      d.HShift = int( i64 )
    case "toc":
      if copt.Args[ 0 ].Data == "true" {
        d.BoolOptions[ "toc" ] = true
      } else { d.BoolOptions[ "toc" ] = false }
    case "mathJax":
      d.BoolOptions[ "mathJax" ] = true
    default:
    }
  }
}

func (scan *Scanner) getTags() (tags []string) {
  scan.Mode = ScanIdents | ScanFloats | ScanStrings | ScanRawStrings
  for scan.Scan().Type != EOF {
    switch scan.LastToken.Type {
    case ScanIdent, ScanInt, ScanHexadecimal, ScanOctal, ScanFloat, ScanString, ScanRawString:
      tags = append( tags, scan.LastToken.Text )
    case ',', ';':
      scan.Error( `Setup (Tags): empty tag` )
      continue
    default: scan.Error( `Setup (Tags): "` + scan.LastToken.Text + `" is not a valid tag` )
    }

    scan.Scan()
    switch scan.LastToken.Type {
    case ',', ';':
    case EOF: return
    default: scan.Error( `Setup (Tags): instert a "," or ";" between tags` )
    }
  }

  return
}
