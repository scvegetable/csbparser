require("functions")
require("json")

local function format(fmt, ...)
    local args = { ... }
    return fmt.gsub(fmt, "{([%d+])}", function (str)
        return tostring(args[tonumber(str)])
    end)
end

local file = io.open("CSParseBinary.json", "rb")
if not file then
    error("can not open file `CSParseBinary.json` for read")
end
local data = file:read("*a")
file:close()

data = json.decode(data)

local codeStr = [[
package main

import (
    "csbparser"
    "fmt"
    flatbuffers"github.com/google/flatbuffers/go"
)

type Decoder interface {
    GetData() interface{}
}

func ConvertOptions(options *csbparser.Options, classname string) Decoder {
    tab := options.Data(nil).Table()
    switch classname {
    case "Node":
        return &WidgetOptions{_tab: tab}
    case "Widget":
        return &WidgetOptions{_tab: tab}
    case "SingleNode":
        return &SingleNodeOptions{_tab: tab}
    case "Sprite":
        return &SpriteOptions{_tab: tab}
    case "Particle":
        return &ParticleSystemOptions{_tab: tab}
    case "GameMap":
        return &GameMapOptions{_tab: tab}
    case "TextButton":
        return &ButtonOptions{_tab: tab}
    case "Button":
        return &ButtonOptions{_tab: tab}
    case "CheckBox":
        return &CheckBoxOptions{_tab: tab}
    case "ImageView":
        return &ImageViewOptions{_tab: tab}
    case "LabelAtlas":
        return &TextAtlasOptions{_tab: tab}
    case "TextAtlas":
        return &TextAtlasOptions{_tab: tab}
    case "LabelBMFont":
        return &TextBMFontOptions{_tab: tab}
    case "TextBMFont":
        return &TextBMFontOptions{_tab: tab}
    case "TextArea":
        return &TextOptions{_tab: tab}
    case "Label":
        return &TextOptions{_tab: tab}
    case "Text":
        return &TextOptions{_tab: tab}
    case "TextField":
        return &TextFieldOptions{_tab: tab}
    case "LoadingBar":
        return &LoadingBarOptions{_tab: tab}
    case "Slider":
        return &SliderOptions{_tab: tab}
    case "Panel":
        return &PanelOptions{_tab: tab}
    case "ScrollView":
        return &ScrollViewOptions{_tab: tab}
    case "PageView":
        return &PageViewOptions{_tab: tab}
    case "ListView":
        return &ListViewOptions{_tab: tab}
    case "ProjectNode":
        return &ProjectNodeOptions{_tab: tab}
    case "Component":
        return &ComponentOptions{_tab: tab}
    case "ComAudio":
        return &ComAudioOptions{_tab:tab}
    }
    fmt.Println("unknown class name", classname)
    return nil
}
]]

function getFuncName(name)
    return string.upper(name:sub(1, 1))..name:sub(2)
end

function genTable(tabDef)
    local code = format([[
type {1} struct { _tab flatbuffers.Table }
func (this *{1})GetData() interface{} {
    table := &csbparser.{1}{}
    table.Init(this._tab.Bytes, this._tab.Pos)
    ret := map[string]interface{}{}
]], tabDef.name, tabDef.name)
    for _, fieldDef in ipairs(tabDef.fields) do
        if fieldDef.base_type == "struct" then
            if tabDef.name == "NodeTree" and fieldDef.name == "options" then
                code = code..format([[
    if options := table.Options(nil); options != nil {
        ret["options"] = ConvertOptions(options, table.Classname()).GetData()
    }
]])
            else
                code = code..format([[
    if {1} := table.{3}(nil); {1} != nil {
        ret["{1}"] = (&{2}{_tab: {1}.Table()}).GetData()
    }
]], fieldDef.name, fieldDef.elem, getFuncName(fieldDef.name))
            end
        elseif fieldDef.base_type == "enum" then
        elseif fieldDef.base_type == "vector" then
            if fieldDef.struct then
                code = code..format([[
    {1}Len := table.{2}Length()
    {1} := []interface{}{}
    for i := 0; i < {1}Len; i++ {
        value := &csbparser.{3}{}
        if table.{2}(value, i) {
            {1} = append({1}, (&{3}{_tab: value.Table()}).GetData())
        }
    }
    ret["{1}"] = {1}
]], fieldDef.name, getFuncName(fieldDef.name), fieldDef.struct)
            elseif fieldDef.enum then
            else
                code = code..format([[
    {1}Len := table.{2}Length()
    {1} := []{3}{}
    for i := 0; i < {1}Len; i++ {
        {1} = append({1}, table.{2}(i))
    }
    ret["{1}"] = {1}
]], fieldDef.name, getFuncName(fieldDef.name), fieldDef.elem)
            end
        else
            code = code..format([[
    ret["{1}"] = table.{2}()
]], fieldDef.name, getFuncName(fieldDef.name))
        end
    end
    code = code.."    return ret\n}\n"
    return code
end

for _, tabDef in ipairs(data.tables) do
    codeStr = codeStr .. genTable(tabDef)
end

local file = io.open("csb_to_json.go", "wb")
if not file then
    error("can not open file `csb_to_json.go` for write")
end
file:write(codeStr.."\n")
file:close()