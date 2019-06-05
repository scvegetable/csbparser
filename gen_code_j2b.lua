require("functions")
require("json")

local function format(fmt, ...)
    local args = { ... }
    return fmt.gsub(fmt, "{([%d+])}", function (str)
        return tostring(args[tonumber(str)])
    end)
end

local mapType = {
    ubyte   = "byte",
    int     = "int32",
    uint    = "uint32",
    float   = "float32",
}
local function getTypeName(baseType)
    return mapType[baseType] or baseType
end

local file = io.open("CSParseBinary.json", "rb")
if not file then
    error("can not open file `CSParseBinary.json` for read")
end
local data = file:read("*a")
file:close()

data = json.decode(data)

local function getStructDef(name)
    for _, tabDef in ipairs(data.tables) do
        if tabDef.name == name then
            return tabDef
        end
    end
end

local codeStr = [[
package main

import (
    "csbparser"
    "fmt"
    flatbuffers"github.com/google/flatbuffers/go"
)

func createOptionsData(builder *flatbuffers.Builder, jValue map[string]interface{}, classname string) (flatbuffers.UOffsetT, bool) {
    switch classname {
    case "Node":
        return CreateWidgetOptions(builder, jValue), true
    case "Widget":
        return CreateWidgetOptions(builder, jValue), true
    case "SingleNode":
        return CreateSingleNodeOptions(builder, jValue), true
    case "Sprite":
        return CreateSpriteOptions(builder, jValue), true
    case "Particle":
        return CreateParticleSystemOptions(builder, jValue), true
    case "GameMap":
        return CreateGameMapOptions(builder, jValue), true
    case "TextButton":
        return CreateButtonOptions(builder, jValue), true
    case "Button":
        return CreateButtonOptions(builder, jValue), true
    case "CheckBox":
        return CreateCheckBoxOptions(builder, jValue), true
    case "ImageView":
        return CreateImageViewOptions(builder, jValue), true
    case "LabelAtlas":
        return CreateTextAtlasOptions(builder, jValue), true
    case "TextAtlas":
        return CreateTextAtlasOptions(builder, jValue), true
    case "LabelBMFont":
        return CreateTextBMFontOptions(builder, jValue), true
    case "TextBMFont":
        return CreateTextBMFontOptions(builder, jValue), true
    case "TextArea":
        return CreateTextOptions(builder, jValue), true
    case "Label":
        return CreateTextOptions(builder, jValue), true
    case "Text":
        return CreateTextOptions(builder, jValue), true
    case "TextField":
        return CreateTextFieldOptions(builder, jValue), true
    case "LoadingBar":
        return CreateLoadingBarOptions(builder, jValue), true
    case "Slider":
        return CreateSliderOptions(builder, jValue), true
    case "Panel":
        return CreatePanelOptions(builder, jValue), true
    case "ScrollView":
        return CreateScrollViewOptions(builder, jValue), true
    case "PageView":
        return CreatePageViewOptions(builder, jValue), true
    case "ListView":
        return CreateListViewOptions(builder, jValue), true
    case "ProjectNode":
        return CreateProjectNodeOptions(builder, jValue), true
    case "Component":
        return CreateComponentOptions(builder, jValue), true
    case "ComAudio":
        return CreateComAudioOptions(builder, jValue), true
    }
    fmt.Println("unknown class name", classname)
    return 0, false
}
func CreateOptionsData(builder *flatbuffers.Builder, jValue map[string]interface{}, classname string) (flatbuffers.UOffsetT, bool) {
    data, ok := createOptionsData(builder, jValue, classname)
    if !ok {
        return data, ok
    }
    csbparser.OptionsStart(builder)
    csbparser.OptionsAddData(builder, data)
    return csbparser.OptionsEnd(builder), true
}
]]

function getFuncName(name)
    return string.upper(name:sub(1, 1))..name:sub(2)
end

function genTable(tabDef)
    local code = format([[
func Create{1}(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
]], tabDef.name)
    for _, fieldDef in ipairs(tabDef.fields) do
        if fieldDef.base_type == "struct" then
            if tabDef.name == "NodeTree" and fieldDef.name == "options" then
                code = code..format([[
    var options flatbuffers.UOffsetT
    hasOptions := false
    if __value, ok := jValue["options"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            if class, ok := jValue["classname"]; ok {
                if class_, ok := class.(string); ok {
                    options, hasOptions = CreateOptionsData(builder, __value_, class_)
                }
            }
        }
    }
]])
            else
                local _tabDef = getStructDef(fieldDef.elem)
                if _tabDef and _tabDef.fixed then
                    code = code..format([[
    var {1} map[string]interface{}
    has{2} := false
    if __value, ok := jValue["{1}"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            {1} = __value_
            has{2} = true
        }
    }
]], fieldDef.name, getFuncName(fieldDef.name))
                else
                    code = code..format([[
    var {1} flatbuffers.UOffsetT
    has{2} := false
    if __value, ok := jValue["{1}"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            {1} = Create{3}(builder, __value_)
            has{2} = true
        }
    }
]], fieldDef.name, getFuncName(fieldDef.name), fieldDef.elem)
                end
            end
        elseif fieldDef.base_type == "enum" then
        elseif fieldDef.base_type == "vector" then
            code = code..format([[
    var {1} flatbuffers.UOffsetT
    has{2} := false
]], fieldDef.name, getFuncName(fieldDef.name))
            if fieldDef.struct then
                code = code..format([[
    if __value, ok := jValue["{1}"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(map[string]interface{}); ok {
                    vector = append(vector, Create{3}(builder, v_))
                }
            }
            csbparser.{4}Start{2}Vector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            {1} = builder.EndVector(len(vector))
            has{2} = true
        }
    }
]], fieldDef.name, getFuncName(fieldDef.name), fieldDef.struct, tabDef.name)
            elseif fieldDef.enum then
            elseif fieldDef.elem == "string" then
                code = code..format([[
    if __value, ok := jValue["{1}"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(string); ok {
                    vector = append(vector, builder.CreateString(v_))
                }
            }
            csbparser.{3}Start{2}Vector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            {1} = builder.EndVector(len(vector))
            has{2} = true
        }
    }
]], fieldDef.name, getFuncName(fieldDef.name), tabDef.name)
            else
                code = code..format([[
    if __value, ok := jValue["{1}"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []{3}{}
            for _, v := range __value_ {
                if v_, ok := v.({6}); ok {
                    {1} = append({1}, {3}(v_))
                }
            }
            csbparser.{5}Start{2}Vector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.Prepend{4}(vector[i])
            }
            {1} = builder.EndVector(len(vector))
            has{2} = true
        }
    }
]], fieldDef.name, getFuncName(fieldDef.name), getTypeName(fieldDef.elem), getFuncName(getTypeName(fieldDef.elem)), tabDef.name, fieldDef.base_type == "bool" and "bool" or "float64")
            end
        elseif fieldDef.base_type == "string" then
            code = code..format([[
    var {1} flatbuffers.UOffsetT
    has{2} := false
    if __value, ok := jValue["{1}"]; ok {
        if __value_, ok := __value.(string); ok {
            {1} = builder.CreateString(__value_)
            has{2} = true
        }
    }
]], fieldDef.name, getFuncName(fieldDef.name))
        else
            code = code..format([[
    var {1} {3}
    has{2} := false
    if __value, ok := jValue["{1}"]; ok {
        if __value_, ok := __value.({4}); ok {
            {1} = {3}(__value_)
            has{2} = true
        }
    }
]], fieldDef.name, getFuncName(fieldDef.name), getTypeName(fieldDef.base_type), fieldDef.base_type == "bool" and "bool" or "float64")
        end
    end

    if tabDef.fixed then
        code = code..format([[
    if]])
        for k, fieldDef in ipairs(tabDef.fields) do
            code = code..format([[ {1}!has{2}]], k > 1 and "&& " or "", getFuncName(fieldDef.name))
        end
        code = code..format([[ {
        return 0
    }
]])
        code = code..format([[
    return csbparser.Create{1}(builder]], tabDef.name)
        for _, fieldDef in ipairs(tabDef.fields) do
            code = code..format([[, {1}]], fieldDef.name)
        end
        code = code..[[)
}
]]
    else
        code = code..format([[
    csbparser.{1}Start(builder)
]], tabDef.name)
        for _, fieldDef in ipairs(tabDef.fields) do
            if fieldDef.base_type == "struct" then
                local _tabDef = getStructDef(fieldDef.elem)
                if _tabDef and _tabDef.fixed then
                    code = code..format([[
    if has{3} {
        offset := Create{4}(builder, {2})
        csbparser.{1}Add{3}(builder, offset)
    }
]], tabDef.name, fieldDef.name, getFuncName(fieldDef.name), fieldDef.elem)
                else
                    code = code..format([[
    if has{3} {
        csbparser.{1}Add{3}(builder, {2})
    }
]], tabDef.name, fieldDef.name, getFuncName(fieldDef.name))
                end
            else
                code = code..format([[
    if has{3} {
        csbparser.{1}Add{3}(builder, {2})
    }
]], tabDef.name, fieldDef.name, getFuncName(fieldDef.name))
            end
        end
        code = code..format([[
    return csbparser.{1}End(builder)
}
]], tabDef.name)
    end
    return code
end

for _, tabDef in ipairs(data.tables) do
    codeStr = codeStr .. genTable(tabDef)
end

local file = io.open("json_to_csb.go", "wb")
if not file then
    error("can not open file `json_to_csb.go` for write")
end
file:write(codeStr.."\n")
file:close()