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
func CreateCSParseBinary(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var version flatbuffers.UOffsetT
    hasVersion := false
    if __value, ok := jValue["version"]; ok {
        if __value_, ok := __value.(string); ok {
            version = builder.CreateString(__value_)
            hasVersion = true
        }
    }
    var textures flatbuffers.UOffsetT
    hasTextures := false
    if __value, ok := jValue["textures"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(string); ok {
                    vector = append(vector, builder.CreateString(v_))
                }
            }
            csbparser.CSParseBinaryStartTexturesVector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            textures = builder.EndVector(len(vector))
            hasTextures = true
        }
    }
    var texturePngs flatbuffers.UOffsetT
    hasTexturePngs := false
    if __value, ok := jValue["texturePngs"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(string); ok {
                    vector = append(vector, builder.CreateString(v_))
                }
            }
            csbparser.CSParseBinaryStartTexturePngsVector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            texturePngs = builder.EndVector(len(vector))
            hasTexturePngs = true
        }
    }
    var nodeTree flatbuffers.UOffsetT
    hasNodeTree := false
    if __value, ok := jValue["nodeTree"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            nodeTree = CreateNodeTree(builder, __value_)
            hasNodeTree = true
        }
    }
    var action flatbuffers.UOffsetT
    hasAction := false
    if __value, ok := jValue["action"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            action = CreateNodeAction(builder, __value_)
            hasAction = true
        }
    }
    var animationList flatbuffers.UOffsetT
    hasAnimationList := false
    if __value, ok := jValue["animationList"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(map[string]interface{}); ok {
                    vector = append(vector, CreateAnimationInfo(builder, v_))
                }
            }
            csbparser.CSParseBinaryStartAnimationListVector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            animationList = builder.EndVector(len(vector))
            hasAnimationList = true
        }
    }
    csbparser.CSParseBinaryStart(builder)
    if hasVersion {
        csbparser.CSParseBinaryAddVersion(builder, version)
    }
    if hasTextures {
        csbparser.CSParseBinaryAddTextures(builder, textures)
    }
    if hasTexturePngs {
        csbparser.CSParseBinaryAddTexturePngs(builder, texturePngs)
    }
    if hasNodeTree {
        csbparser.CSParseBinaryAddNodeTree(builder, nodeTree)
    }
    if hasAction {
        csbparser.CSParseBinaryAddAction(builder, action)
    }
    if hasAnimationList {
        csbparser.CSParseBinaryAddAnimationList(builder, animationList)
    }
    return csbparser.CSParseBinaryEnd(builder)
}
func CreateNodeTree(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var classname flatbuffers.UOffsetT
    hasClassname := false
    if __value, ok := jValue["classname"]; ok {
        if __value_, ok := __value.(string); ok {
            classname = builder.CreateString(__value_)
            hasClassname = true
        }
    }
    var children flatbuffers.UOffsetT
    hasChildren := false
    if __value, ok := jValue["children"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(map[string]interface{}); ok {
                    vector = append(vector, CreateNodeTree(builder, v_))
                }
            }
            csbparser.NodeTreeStartChildrenVector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            children = builder.EndVector(len(vector))
            hasChildren = true
        }
    }
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
    var customClassName flatbuffers.UOffsetT
    hasCustomClassName := false
    if __value, ok := jValue["customClassName"]; ok {
        if __value_, ok := __value.(string); ok {
            customClassName = builder.CreateString(__value_)
            hasCustomClassName = true
        }
    }
    csbparser.NodeTreeStart(builder)
    if hasClassname {
        csbparser.NodeTreeAddClassname(builder, classname)
    }
    if hasChildren {
        csbparser.NodeTreeAddChildren(builder, children)
    }
    if hasOptions {
        csbparser.NodeTreeAddOptions(builder, options)
    }
    if hasCustomClassName {
        csbparser.NodeTreeAddCustomClassName(builder, customClassName)
    }
    return csbparser.NodeTreeEnd(builder)
}
func CreateOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var data flatbuffers.UOffsetT
    hasData := false
    if __value, ok := jValue["data"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            data = CreateWidgetOptions(builder, __value_)
            hasData = true
        }
    }
    csbparser.OptionsStart(builder)
    if hasData {
        csbparser.OptionsAddData(builder, data)
    }
    return csbparser.OptionsEnd(builder)
}
func CreateWidgetOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var name flatbuffers.UOffsetT
    hasName := false
    if __value, ok := jValue["name"]; ok {
        if __value_, ok := __value.(string); ok {
            name = builder.CreateString(__value_)
            hasName = true
        }
    }
    var actionTag int32
    hasActionTag := false
    if __value, ok := jValue["actionTag"]; ok {
        if __value_, ok := __value.(float64); ok {
            actionTag = int32(__value_)
            hasActionTag = true
        }
    }
    var rotationSkew map[string]interface{}
    hasRotationSkew := false
    if __value, ok := jValue["rotationSkew"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            rotationSkew = __value_
            hasRotationSkew = true
        }
    }
    var zOrder int32
    hasZOrder := false
    if __value, ok := jValue["zOrder"]; ok {
        if __value_, ok := __value.(float64); ok {
            zOrder = int32(__value_)
            hasZOrder = true
        }
    }
    var visible bool
    hasVisible := false
    if __value, ok := jValue["visible"]; ok {
        if __value_, ok := __value.(bool); ok {
            visible = bool(__value_)
            hasVisible = true
        }
    }
    var alpha byte
    hasAlpha := false
    if __value, ok := jValue["alpha"]; ok {
        if __value_, ok := __value.(float64); ok {
            alpha = byte(__value_)
            hasAlpha = true
        }
    }
    var tag int32
    hasTag := false
    if __value, ok := jValue["tag"]; ok {
        if __value_, ok := __value.(float64); ok {
            tag = int32(__value_)
            hasTag = true
        }
    }
    var position map[string]interface{}
    hasPosition := false
    if __value, ok := jValue["position"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            position = __value_
            hasPosition = true
        }
    }
    var scale map[string]interface{}
    hasScale := false
    if __value, ok := jValue["scale"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scale = __value_
            hasScale = true
        }
    }
    var anchorPoint map[string]interface{}
    hasAnchorPoint := false
    if __value, ok := jValue["anchorPoint"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            anchorPoint = __value_
            hasAnchorPoint = true
        }
    }
    var color map[string]interface{}
    hasColor := false
    if __value, ok := jValue["color"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            color = __value_
            hasColor = true
        }
    }
    var size map[string]interface{}
    hasSize := false
    if __value, ok := jValue["size"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            size = __value_
            hasSize = true
        }
    }
    var flipX bool
    hasFlipX := false
    if __value, ok := jValue["flipX"]; ok {
        if __value_, ok := __value.(bool); ok {
            flipX = bool(__value_)
            hasFlipX = true
        }
    }
    var flipY bool
    hasFlipY := false
    if __value, ok := jValue["flipY"]; ok {
        if __value_, ok := __value.(bool); ok {
            flipY = bool(__value_)
            hasFlipY = true
        }
    }
    var ignoreSize bool
    hasIgnoreSize := false
    if __value, ok := jValue["ignoreSize"]; ok {
        if __value_, ok := __value.(bool); ok {
            ignoreSize = bool(__value_)
            hasIgnoreSize = true
        }
    }
    var touchEnabled bool
    hasTouchEnabled := false
    if __value, ok := jValue["touchEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            touchEnabled = bool(__value_)
            hasTouchEnabled = true
        }
    }
    var frameEvent flatbuffers.UOffsetT
    hasFrameEvent := false
    if __value, ok := jValue["frameEvent"]; ok {
        if __value_, ok := __value.(string); ok {
            frameEvent = builder.CreateString(__value_)
            hasFrameEvent = true
        }
    }
    var customProperty flatbuffers.UOffsetT
    hasCustomProperty := false
    if __value, ok := jValue["customProperty"]; ok {
        if __value_, ok := __value.(string); ok {
            customProperty = builder.CreateString(__value_)
            hasCustomProperty = true
        }
    }
    var callBackType flatbuffers.UOffsetT
    hasCallBackType := false
    if __value, ok := jValue["callBackType"]; ok {
        if __value_, ok := __value.(string); ok {
            callBackType = builder.CreateString(__value_)
            hasCallBackType = true
        }
    }
    var callBackName flatbuffers.UOffsetT
    hasCallBackName := false
    if __value, ok := jValue["callBackName"]; ok {
        if __value_, ok := __value.(string); ok {
            callBackName = builder.CreateString(__value_)
            hasCallBackName = true
        }
    }
    var layoutComponent flatbuffers.UOffsetT
    hasLayoutComponent := false
    if __value, ok := jValue["layoutComponent"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            layoutComponent = CreateLayoutComponentTable(builder, __value_)
            hasLayoutComponent = true
        }
    }
    csbparser.WidgetOptionsStart(builder)
    if hasName {
        csbparser.WidgetOptionsAddName(builder, name)
    }
    if hasActionTag {
        csbparser.WidgetOptionsAddActionTag(builder, actionTag)
    }
    if hasRotationSkew {
        offset := CreateRotationSkew(builder, rotationSkew)
        csbparser.WidgetOptionsAddRotationSkew(builder, offset)
    }
    if hasZOrder {
        csbparser.WidgetOptionsAddZOrder(builder, zOrder)
    }
    if hasVisible {
        csbparser.WidgetOptionsAddVisible(builder, visible)
    }
    if hasAlpha {
        csbparser.WidgetOptionsAddAlpha(builder, alpha)
    }
    if hasTag {
        csbparser.WidgetOptionsAddTag(builder, tag)
    }
    if hasPosition {
        offset := CreatePosition(builder, position)
        csbparser.WidgetOptionsAddPosition(builder, offset)
    }
    if hasScale {
        offset := CreateScale(builder, scale)
        csbparser.WidgetOptionsAddScale(builder, offset)
    }
    if hasAnchorPoint {
        offset := CreateAnchorPoint(builder, anchorPoint)
        csbparser.WidgetOptionsAddAnchorPoint(builder, offset)
    }
    if hasColor {
        offset := CreateColor(builder, color)
        csbparser.WidgetOptionsAddColor(builder, offset)
    }
    if hasSize {
        offset := CreateFlatSize(builder, size)
        csbparser.WidgetOptionsAddSize(builder, offset)
    }
    if hasFlipX {
        csbparser.WidgetOptionsAddFlipX(builder, flipX)
    }
    if hasFlipY {
        csbparser.WidgetOptionsAddFlipY(builder, flipY)
    }
    if hasIgnoreSize {
        csbparser.WidgetOptionsAddIgnoreSize(builder, ignoreSize)
    }
    if hasTouchEnabled {
        csbparser.WidgetOptionsAddTouchEnabled(builder, touchEnabled)
    }
    if hasFrameEvent {
        csbparser.WidgetOptionsAddFrameEvent(builder, frameEvent)
    }
    if hasCustomProperty {
        csbparser.WidgetOptionsAddCustomProperty(builder, customProperty)
    }
    if hasCallBackType {
        csbparser.WidgetOptionsAddCallBackType(builder, callBackType)
    }
    if hasCallBackName {
        csbparser.WidgetOptionsAddCallBackName(builder, callBackName)
    }
    if hasLayoutComponent {
        csbparser.WidgetOptionsAddLayoutComponent(builder, layoutComponent)
    }
    return csbparser.WidgetOptionsEnd(builder)
}
func CreateLayoutComponentTable(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var positionXPercentEnabled bool
    hasPositionXPercentEnabled := false
    if __value, ok := jValue["positionXPercentEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            positionXPercentEnabled = bool(__value_)
            hasPositionXPercentEnabled = true
        }
    }
    var positionYPercentEnabled bool
    hasPositionYPercentEnabled := false
    if __value, ok := jValue["positionYPercentEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            positionYPercentEnabled = bool(__value_)
            hasPositionYPercentEnabled = true
        }
    }
    var positionXPercent float32
    hasPositionXPercent := false
    if __value, ok := jValue["positionXPercent"]; ok {
        if __value_, ok := __value.(float64); ok {
            positionXPercent = float32(__value_)
            hasPositionXPercent = true
        }
    }
    var positionYPercent float32
    hasPositionYPercent := false
    if __value, ok := jValue["positionYPercent"]; ok {
        if __value_, ok := __value.(float64); ok {
            positionYPercent = float32(__value_)
            hasPositionYPercent = true
        }
    }
    var sizeXPercentEnable bool
    hasSizeXPercentEnable := false
    if __value, ok := jValue["sizeXPercentEnable"]; ok {
        if __value_, ok := __value.(bool); ok {
            sizeXPercentEnable = bool(__value_)
            hasSizeXPercentEnable = true
        }
    }
    var sizeYPercentEnable bool
    hasSizeYPercentEnable := false
    if __value, ok := jValue["sizeYPercentEnable"]; ok {
        if __value_, ok := __value.(bool); ok {
            sizeYPercentEnable = bool(__value_)
            hasSizeYPercentEnable = true
        }
    }
    var sizeXPercent float32
    hasSizeXPercent := false
    if __value, ok := jValue["sizeXPercent"]; ok {
        if __value_, ok := __value.(float64); ok {
            sizeXPercent = float32(__value_)
            hasSizeXPercent = true
        }
    }
    var sizeYPercent float32
    hasSizeYPercent := false
    if __value, ok := jValue["sizeYPercent"]; ok {
        if __value_, ok := __value.(float64); ok {
            sizeYPercent = float32(__value_)
            hasSizeYPercent = true
        }
    }
    var stretchHorizontalEnabled bool
    hasStretchHorizontalEnabled := false
    if __value, ok := jValue["stretchHorizontalEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            stretchHorizontalEnabled = bool(__value_)
            hasStretchHorizontalEnabled = true
        }
    }
    var stretchVerticalEnabled bool
    hasStretchVerticalEnabled := false
    if __value, ok := jValue["stretchVerticalEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            stretchVerticalEnabled = bool(__value_)
            hasStretchVerticalEnabled = true
        }
    }
    var horizontalEdge flatbuffers.UOffsetT
    hasHorizontalEdge := false
    if __value, ok := jValue["horizontalEdge"]; ok {
        if __value_, ok := __value.(string); ok {
            horizontalEdge = builder.CreateString(__value_)
            hasHorizontalEdge = true
        }
    }
    var verticalEdge flatbuffers.UOffsetT
    hasVerticalEdge := false
    if __value, ok := jValue["verticalEdge"]; ok {
        if __value_, ok := __value.(string); ok {
            verticalEdge = builder.CreateString(__value_)
            hasVerticalEdge = true
        }
    }
    var leftMargin float32
    hasLeftMargin := false
    if __value, ok := jValue["leftMargin"]; ok {
        if __value_, ok := __value.(float64); ok {
            leftMargin = float32(__value_)
            hasLeftMargin = true
        }
    }
    var rightMargin float32
    hasRightMargin := false
    if __value, ok := jValue["rightMargin"]; ok {
        if __value_, ok := __value.(float64); ok {
            rightMargin = float32(__value_)
            hasRightMargin = true
        }
    }
    var topMargin float32
    hasTopMargin := false
    if __value, ok := jValue["topMargin"]; ok {
        if __value_, ok := __value.(float64); ok {
            topMargin = float32(__value_)
            hasTopMargin = true
        }
    }
    var bottomMargin float32
    hasBottomMargin := false
    if __value, ok := jValue["bottomMargin"]; ok {
        if __value_, ok := __value.(float64); ok {
            bottomMargin = float32(__value_)
            hasBottomMargin = true
        }
    }
    csbparser.LayoutComponentTableStart(builder)
    if hasPositionXPercentEnabled {
        csbparser.LayoutComponentTableAddPositionXPercentEnabled(builder, positionXPercentEnabled)
    }
    if hasPositionYPercentEnabled {
        csbparser.LayoutComponentTableAddPositionYPercentEnabled(builder, positionYPercentEnabled)
    }
    if hasPositionXPercent {
        csbparser.LayoutComponentTableAddPositionXPercent(builder, positionXPercent)
    }
    if hasPositionYPercent {
        csbparser.LayoutComponentTableAddPositionYPercent(builder, positionYPercent)
    }
    if hasSizeXPercentEnable {
        csbparser.LayoutComponentTableAddSizeXPercentEnable(builder, sizeXPercentEnable)
    }
    if hasSizeYPercentEnable {
        csbparser.LayoutComponentTableAddSizeYPercentEnable(builder, sizeYPercentEnable)
    }
    if hasSizeXPercent {
        csbparser.LayoutComponentTableAddSizeXPercent(builder, sizeXPercent)
    }
    if hasSizeYPercent {
        csbparser.LayoutComponentTableAddSizeYPercent(builder, sizeYPercent)
    }
    if hasStretchHorizontalEnabled {
        csbparser.LayoutComponentTableAddStretchHorizontalEnabled(builder, stretchHorizontalEnabled)
    }
    if hasStretchVerticalEnabled {
        csbparser.LayoutComponentTableAddStretchVerticalEnabled(builder, stretchVerticalEnabled)
    }
    if hasHorizontalEdge {
        csbparser.LayoutComponentTableAddHorizontalEdge(builder, horizontalEdge)
    }
    if hasVerticalEdge {
        csbparser.LayoutComponentTableAddVerticalEdge(builder, verticalEdge)
    }
    if hasLeftMargin {
        csbparser.LayoutComponentTableAddLeftMargin(builder, leftMargin)
    }
    if hasRightMargin {
        csbparser.LayoutComponentTableAddRightMargin(builder, rightMargin)
    }
    if hasTopMargin {
        csbparser.LayoutComponentTableAddTopMargin(builder, topMargin)
    }
    if hasBottomMargin {
        csbparser.LayoutComponentTableAddBottomMargin(builder, bottomMargin)
    }
    return csbparser.LayoutComponentTableEnd(builder)
}
func CreateSingleNodeOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var nodeOptions flatbuffers.UOffsetT
    hasNodeOptions := false
    if __value, ok := jValue["nodeOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            nodeOptions = CreateWidgetOptions(builder, __value_)
            hasNodeOptions = true
        }
    }
    csbparser.SingleNodeOptionsStart(builder)
    if hasNodeOptions {
        csbparser.SingleNodeOptionsAddNodeOptions(builder, nodeOptions)
    }
    return csbparser.SingleNodeOptionsEnd(builder)
}
func CreateSpriteOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var nodeOptions flatbuffers.UOffsetT
    hasNodeOptions := false
    if __value, ok := jValue["nodeOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            nodeOptions = CreateWidgetOptions(builder, __value_)
            hasNodeOptions = true
        }
    }
    var fileNameData flatbuffers.UOffsetT
    hasFileNameData := false
    if __value, ok := jValue["fileNameData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fileNameData = CreateResourceData(builder, __value_)
            hasFileNameData = true
        }
    }
    var blendFunc map[string]interface{}
    hasBlendFunc := false
    if __value, ok := jValue["blendFunc"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            blendFunc = __value_
            hasBlendFunc = true
        }
    }
    csbparser.SpriteOptionsStart(builder)
    if hasNodeOptions {
        csbparser.SpriteOptionsAddNodeOptions(builder, nodeOptions)
    }
    if hasFileNameData {
        csbparser.SpriteOptionsAddFileNameData(builder, fileNameData)
    }
    if hasBlendFunc {
        offset := CreateBlendFunc(builder, blendFunc)
        csbparser.SpriteOptionsAddBlendFunc(builder, offset)
    }
    return csbparser.SpriteOptionsEnd(builder)
}
func CreateParticleSystemOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var nodeOptions flatbuffers.UOffsetT
    hasNodeOptions := false
    if __value, ok := jValue["nodeOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            nodeOptions = CreateWidgetOptions(builder, __value_)
            hasNodeOptions = true
        }
    }
    var fileNameData flatbuffers.UOffsetT
    hasFileNameData := false
    if __value, ok := jValue["fileNameData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fileNameData = CreateResourceData(builder, __value_)
            hasFileNameData = true
        }
    }
    var blendFunc map[string]interface{}
    hasBlendFunc := false
    if __value, ok := jValue["blendFunc"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            blendFunc = __value_
            hasBlendFunc = true
        }
    }
    csbparser.ParticleSystemOptionsStart(builder)
    if hasNodeOptions {
        csbparser.ParticleSystemOptionsAddNodeOptions(builder, nodeOptions)
    }
    if hasFileNameData {
        csbparser.ParticleSystemOptionsAddFileNameData(builder, fileNameData)
    }
    if hasBlendFunc {
        offset := CreateBlendFunc(builder, blendFunc)
        csbparser.ParticleSystemOptionsAddBlendFunc(builder, offset)
    }
    return csbparser.ParticleSystemOptionsEnd(builder)
}
func CreateGameMapOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var nodeOptions flatbuffers.UOffsetT
    hasNodeOptions := false
    if __value, ok := jValue["nodeOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            nodeOptions = CreateWidgetOptions(builder, __value_)
            hasNodeOptions = true
        }
    }
    var fileNameData flatbuffers.UOffsetT
    hasFileNameData := false
    if __value, ok := jValue["fileNameData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fileNameData = CreateResourceData(builder, __value_)
            hasFileNameData = true
        }
    }
    csbparser.GameMapOptionsStart(builder)
    if hasNodeOptions {
        csbparser.GameMapOptionsAddNodeOptions(builder, nodeOptions)
    }
    if hasFileNameData {
        csbparser.GameMapOptionsAddFileNameData(builder, fileNameData)
    }
    return csbparser.GameMapOptionsEnd(builder)
}
func CreateButtonOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var normalData flatbuffers.UOffsetT
    hasNormalData := false
    if __value, ok := jValue["normalData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            normalData = CreateResourceData(builder, __value_)
            hasNormalData = true
        }
    }
    var pressedData flatbuffers.UOffsetT
    hasPressedData := false
    if __value, ok := jValue["pressedData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            pressedData = CreateResourceData(builder, __value_)
            hasPressedData = true
        }
    }
    var disabledData flatbuffers.UOffsetT
    hasDisabledData := false
    if __value, ok := jValue["disabledData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            disabledData = CreateResourceData(builder, __value_)
            hasDisabledData = true
        }
    }
    var fontResource flatbuffers.UOffsetT
    hasFontResource := false
    if __value, ok := jValue["fontResource"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fontResource = CreateResourceData(builder, __value_)
            hasFontResource = true
        }
    }
    var text flatbuffers.UOffsetT
    hasText := false
    if __value, ok := jValue["text"]; ok {
        if __value_, ok := __value.(string); ok {
            text = builder.CreateString(__value_)
            hasText = true
        }
    }
    var fontName flatbuffers.UOffsetT
    hasFontName := false
    if __value, ok := jValue["fontName"]; ok {
        if __value_, ok := __value.(string); ok {
            fontName = builder.CreateString(__value_)
            hasFontName = true
        }
    }
    var fontSize int32
    hasFontSize := false
    if __value, ok := jValue["fontSize"]; ok {
        if __value_, ok := __value.(float64); ok {
            fontSize = int32(__value_)
            hasFontSize = true
        }
    }
    var textColor map[string]interface{}
    hasTextColor := false
    if __value, ok := jValue["textColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            textColor = __value_
            hasTextColor = true
        }
    }
    var capInsets map[string]interface{}
    hasCapInsets := false
    if __value, ok := jValue["capInsets"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            capInsets = __value_
            hasCapInsets = true
        }
    }
    var scale9Size map[string]interface{}
    hasScale9Size := false
    if __value, ok := jValue["scale9Size"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scale9Size = __value_
            hasScale9Size = true
        }
    }
    var scale9Enabled bool
    hasScale9Enabled := false
    if __value, ok := jValue["scale9Enabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            scale9Enabled = bool(__value_)
            hasScale9Enabled = true
        }
    }
    var displaystate bool
    hasDisplaystate := false
    if __value, ok := jValue["displaystate"]; ok {
        if __value_, ok := __value.(bool); ok {
            displaystate = bool(__value_)
            hasDisplaystate = true
        }
    }
    var outlineEnabled bool
    hasOutlineEnabled := false
    if __value, ok := jValue["outlineEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            outlineEnabled = bool(__value_)
            hasOutlineEnabled = true
        }
    }
    var outlineColor map[string]interface{}
    hasOutlineColor := false
    if __value, ok := jValue["outlineColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            outlineColor = __value_
            hasOutlineColor = true
        }
    }
    var outlineSize int32
    hasOutlineSize := false
    if __value, ok := jValue["outlineSize"]; ok {
        if __value_, ok := __value.(float64); ok {
            outlineSize = int32(__value_)
            hasOutlineSize = true
        }
    }
    var shadowEnabled bool
    hasShadowEnabled := false
    if __value, ok := jValue["shadowEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            shadowEnabled = bool(__value_)
            hasShadowEnabled = true
        }
    }
    var shadowColor map[string]interface{}
    hasShadowColor := false
    if __value, ok := jValue["shadowColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            shadowColor = __value_
            hasShadowColor = true
        }
    }
    var shadowOffsetX float32
    hasShadowOffsetX := false
    if __value, ok := jValue["shadowOffsetX"]; ok {
        if __value_, ok := __value.(float64); ok {
            shadowOffsetX = float32(__value_)
            hasShadowOffsetX = true
        }
    }
    var shadowOffsetY float32
    hasShadowOffsetY := false
    if __value, ok := jValue["shadowOffsetY"]; ok {
        if __value_, ok := __value.(float64); ok {
            shadowOffsetY = float32(__value_)
            hasShadowOffsetY = true
        }
    }
    var shadowBlurRadius int32
    hasShadowBlurRadius := false
    if __value, ok := jValue["shadowBlurRadius"]; ok {
        if __value_, ok := __value.(float64); ok {
            shadowBlurRadius = int32(__value_)
            hasShadowBlurRadius = true
        }
    }
    csbparser.ButtonOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.ButtonOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasNormalData {
        csbparser.ButtonOptionsAddNormalData(builder, normalData)
    }
    if hasPressedData {
        csbparser.ButtonOptionsAddPressedData(builder, pressedData)
    }
    if hasDisabledData {
        csbparser.ButtonOptionsAddDisabledData(builder, disabledData)
    }
    if hasFontResource {
        csbparser.ButtonOptionsAddFontResource(builder, fontResource)
    }
    if hasText {
        csbparser.ButtonOptionsAddText(builder, text)
    }
    if hasFontName {
        csbparser.ButtonOptionsAddFontName(builder, fontName)
    }
    if hasFontSize {
        csbparser.ButtonOptionsAddFontSize(builder, fontSize)
    }
    if hasTextColor {
        offset := CreateColor(builder, textColor)
        csbparser.ButtonOptionsAddTextColor(builder, offset)
    }
    if hasCapInsets {
        offset := CreateCapInsets(builder, capInsets)
        csbparser.ButtonOptionsAddCapInsets(builder, offset)
    }
    if hasScale9Size {
        offset := CreateFlatSize(builder, scale9Size)
        csbparser.ButtonOptionsAddScale9Size(builder, offset)
    }
    if hasScale9Enabled {
        csbparser.ButtonOptionsAddScale9Enabled(builder, scale9Enabled)
    }
    if hasDisplaystate {
        csbparser.ButtonOptionsAddDisplaystate(builder, displaystate)
    }
    if hasOutlineEnabled {
        csbparser.ButtonOptionsAddOutlineEnabled(builder, outlineEnabled)
    }
    if hasOutlineColor {
        offset := CreateColor(builder, outlineColor)
        csbparser.ButtonOptionsAddOutlineColor(builder, offset)
    }
    if hasOutlineSize {
        csbparser.ButtonOptionsAddOutlineSize(builder, outlineSize)
    }
    if hasShadowEnabled {
        csbparser.ButtonOptionsAddShadowEnabled(builder, shadowEnabled)
    }
    if hasShadowColor {
        offset := CreateColor(builder, shadowColor)
        csbparser.ButtonOptionsAddShadowColor(builder, offset)
    }
    if hasShadowOffsetX {
        csbparser.ButtonOptionsAddShadowOffsetX(builder, shadowOffsetX)
    }
    if hasShadowOffsetY {
        csbparser.ButtonOptionsAddShadowOffsetY(builder, shadowOffsetY)
    }
    if hasShadowBlurRadius {
        csbparser.ButtonOptionsAddShadowBlurRadius(builder, shadowBlurRadius)
    }
    return csbparser.ButtonOptionsEnd(builder)
}
func CreateCheckBoxOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var backGroundBoxData flatbuffers.UOffsetT
    hasBackGroundBoxData := false
    if __value, ok := jValue["backGroundBoxData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            backGroundBoxData = CreateResourceData(builder, __value_)
            hasBackGroundBoxData = true
        }
    }
    var backGroundBoxSelectedData flatbuffers.UOffsetT
    hasBackGroundBoxSelectedData := false
    if __value, ok := jValue["backGroundBoxSelectedData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            backGroundBoxSelectedData = CreateResourceData(builder, __value_)
            hasBackGroundBoxSelectedData = true
        }
    }
    var frontCrossData flatbuffers.UOffsetT
    hasFrontCrossData := false
    if __value, ok := jValue["frontCrossData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            frontCrossData = CreateResourceData(builder, __value_)
            hasFrontCrossData = true
        }
    }
    var backGroundBoxDisabledData flatbuffers.UOffsetT
    hasBackGroundBoxDisabledData := false
    if __value, ok := jValue["backGroundBoxDisabledData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            backGroundBoxDisabledData = CreateResourceData(builder, __value_)
            hasBackGroundBoxDisabledData = true
        }
    }
    var frontCrossDisabledData flatbuffers.UOffsetT
    hasFrontCrossDisabledData := false
    if __value, ok := jValue["frontCrossDisabledData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            frontCrossDisabledData = CreateResourceData(builder, __value_)
            hasFrontCrossDisabledData = true
        }
    }
    var selectedState bool
    hasSelectedState := false
    if __value, ok := jValue["selectedState"]; ok {
        if __value_, ok := __value.(bool); ok {
            selectedState = bool(__value_)
            hasSelectedState = true
        }
    }
    var displaystate bool
    hasDisplaystate := false
    if __value, ok := jValue["displaystate"]; ok {
        if __value_, ok := __value.(bool); ok {
            displaystate = bool(__value_)
            hasDisplaystate = true
        }
    }
    csbparser.CheckBoxOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.CheckBoxOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasBackGroundBoxData {
        csbparser.CheckBoxOptionsAddBackGroundBoxData(builder, backGroundBoxData)
    }
    if hasBackGroundBoxSelectedData {
        csbparser.CheckBoxOptionsAddBackGroundBoxSelectedData(builder, backGroundBoxSelectedData)
    }
    if hasFrontCrossData {
        csbparser.CheckBoxOptionsAddFrontCrossData(builder, frontCrossData)
    }
    if hasBackGroundBoxDisabledData {
        csbparser.CheckBoxOptionsAddBackGroundBoxDisabledData(builder, backGroundBoxDisabledData)
    }
    if hasFrontCrossDisabledData {
        csbparser.CheckBoxOptionsAddFrontCrossDisabledData(builder, frontCrossDisabledData)
    }
    if hasSelectedState {
        csbparser.CheckBoxOptionsAddSelectedState(builder, selectedState)
    }
    if hasDisplaystate {
        csbparser.CheckBoxOptionsAddDisplaystate(builder, displaystate)
    }
    return csbparser.CheckBoxOptionsEnd(builder)
}
func CreateImageViewOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var fileNameData flatbuffers.UOffsetT
    hasFileNameData := false
    if __value, ok := jValue["fileNameData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fileNameData = CreateResourceData(builder, __value_)
            hasFileNameData = true
        }
    }
    var capInsets map[string]interface{}
    hasCapInsets := false
    if __value, ok := jValue["capInsets"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            capInsets = __value_
            hasCapInsets = true
        }
    }
    var scale9Size map[string]interface{}
    hasScale9Size := false
    if __value, ok := jValue["scale9Size"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scale9Size = __value_
            hasScale9Size = true
        }
    }
    var scale9Enabled bool
    hasScale9Enabled := false
    if __value, ok := jValue["scale9Enabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            scale9Enabled = bool(__value_)
            hasScale9Enabled = true
        }
    }
    csbparser.ImageViewOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.ImageViewOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasFileNameData {
        csbparser.ImageViewOptionsAddFileNameData(builder, fileNameData)
    }
    if hasCapInsets {
        offset := CreateCapInsets(builder, capInsets)
        csbparser.ImageViewOptionsAddCapInsets(builder, offset)
    }
    if hasScale9Size {
        offset := CreateFlatSize(builder, scale9Size)
        csbparser.ImageViewOptionsAddScale9Size(builder, offset)
    }
    if hasScale9Enabled {
        csbparser.ImageViewOptionsAddScale9Enabled(builder, scale9Enabled)
    }
    return csbparser.ImageViewOptionsEnd(builder)
}
func CreateTextAtlasOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var charMapFileData flatbuffers.UOffsetT
    hasCharMapFileData := false
    if __value, ok := jValue["charMapFileData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            charMapFileData = CreateResourceData(builder, __value_)
            hasCharMapFileData = true
        }
    }
    var stringValue flatbuffers.UOffsetT
    hasStringValue := false
    if __value, ok := jValue["stringValue"]; ok {
        if __value_, ok := __value.(string); ok {
            stringValue = builder.CreateString(__value_)
            hasStringValue = true
        }
    }
    var startCharMap flatbuffers.UOffsetT
    hasStartCharMap := false
    if __value, ok := jValue["startCharMap"]; ok {
        if __value_, ok := __value.(string); ok {
            startCharMap = builder.CreateString(__value_)
            hasStartCharMap = true
        }
    }
    var itemWidth int32
    hasItemWidth := false
    if __value, ok := jValue["itemWidth"]; ok {
        if __value_, ok := __value.(float64); ok {
            itemWidth = int32(__value_)
            hasItemWidth = true
        }
    }
    var itemHeight int32
    hasItemHeight := false
    if __value, ok := jValue["itemHeight"]; ok {
        if __value_, ok := __value.(float64); ok {
            itemHeight = int32(__value_)
            hasItemHeight = true
        }
    }
    csbparser.TextAtlasOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.TextAtlasOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasCharMapFileData {
        csbparser.TextAtlasOptionsAddCharMapFileData(builder, charMapFileData)
    }
    if hasStringValue {
        csbparser.TextAtlasOptionsAddStringValue(builder, stringValue)
    }
    if hasStartCharMap {
        csbparser.TextAtlasOptionsAddStartCharMap(builder, startCharMap)
    }
    if hasItemWidth {
        csbparser.TextAtlasOptionsAddItemWidth(builder, itemWidth)
    }
    if hasItemHeight {
        csbparser.TextAtlasOptionsAddItemHeight(builder, itemHeight)
    }
    return csbparser.TextAtlasOptionsEnd(builder)
}
func CreateTextBMFontOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var fileNameData flatbuffers.UOffsetT
    hasFileNameData := false
    if __value, ok := jValue["fileNameData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fileNameData = CreateResourceData(builder, __value_)
            hasFileNameData = true
        }
    }
    var text flatbuffers.UOffsetT
    hasText := false
    if __value, ok := jValue["text"]; ok {
        if __value_, ok := __value.(string); ok {
            text = builder.CreateString(__value_)
            hasText = true
        }
    }
    csbparser.TextBMFontOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.TextBMFontOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasFileNameData {
        csbparser.TextBMFontOptionsAddFileNameData(builder, fileNameData)
    }
    if hasText {
        csbparser.TextBMFontOptionsAddText(builder, text)
    }
    return csbparser.TextBMFontOptionsEnd(builder)
}
func CreateTextOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var fontResource flatbuffers.UOffsetT
    hasFontResource := false
    if __value, ok := jValue["fontResource"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fontResource = CreateResourceData(builder, __value_)
            hasFontResource = true
        }
    }
    var fontName flatbuffers.UOffsetT
    hasFontName := false
    if __value, ok := jValue["fontName"]; ok {
        if __value_, ok := __value.(string); ok {
            fontName = builder.CreateString(__value_)
            hasFontName = true
        }
    }
    var fontSize int32
    hasFontSize := false
    if __value, ok := jValue["fontSize"]; ok {
        if __value_, ok := __value.(float64); ok {
            fontSize = int32(__value_)
            hasFontSize = true
        }
    }
    var text flatbuffers.UOffsetT
    hasText := false
    if __value, ok := jValue["text"]; ok {
        if __value_, ok := __value.(string); ok {
            text = builder.CreateString(__value_)
            hasText = true
        }
    }
    var areaWidth int32
    hasAreaWidth := false
    if __value, ok := jValue["areaWidth"]; ok {
        if __value_, ok := __value.(float64); ok {
            areaWidth = int32(__value_)
            hasAreaWidth = true
        }
    }
    var areaHeight int32
    hasAreaHeight := false
    if __value, ok := jValue["areaHeight"]; ok {
        if __value_, ok := __value.(float64); ok {
            areaHeight = int32(__value_)
            hasAreaHeight = true
        }
    }
    var hAlignment int32
    hasHAlignment := false
    if __value, ok := jValue["hAlignment"]; ok {
        if __value_, ok := __value.(float64); ok {
            hAlignment = int32(__value_)
            hasHAlignment = true
        }
    }
    var vAlignment int32
    hasVAlignment := false
    if __value, ok := jValue["vAlignment"]; ok {
        if __value_, ok := __value.(float64); ok {
            vAlignment = int32(__value_)
            hasVAlignment = true
        }
    }
    var touchScaleEnable bool
    hasTouchScaleEnable := false
    if __value, ok := jValue["touchScaleEnable"]; ok {
        if __value_, ok := __value.(bool); ok {
            touchScaleEnable = bool(__value_)
            hasTouchScaleEnable = true
        }
    }
    var isCustomSize bool
    hasIsCustomSize := false
    if __value, ok := jValue["isCustomSize"]; ok {
        if __value_, ok := __value.(bool); ok {
            isCustomSize = bool(__value_)
            hasIsCustomSize = true
        }
    }
    var outlineEnabled bool
    hasOutlineEnabled := false
    if __value, ok := jValue["outlineEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            outlineEnabled = bool(__value_)
            hasOutlineEnabled = true
        }
    }
    var outlineColor map[string]interface{}
    hasOutlineColor := false
    if __value, ok := jValue["outlineColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            outlineColor = __value_
            hasOutlineColor = true
        }
    }
    var outlineSize int32
    hasOutlineSize := false
    if __value, ok := jValue["outlineSize"]; ok {
        if __value_, ok := __value.(float64); ok {
            outlineSize = int32(__value_)
            hasOutlineSize = true
        }
    }
    var shadowEnabled bool
    hasShadowEnabled := false
    if __value, ok := jValue["shadowEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            shadowEnabled = bool(__value_)
            hasShadowEnabled = true
        }
    }
    var shadowColor map[string]interface{}
    hasShadowColor := false
    if __value, ok := jValue["shadowColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            shadowColor = __value_
            hasShadowColor = true
        }
    }
    var shadowOffsetX float32
    hasShadowOffsetX := false
    if __value, ok := jValue["shadowOffsetX"]; ok {
        if __value_, ok := __value.(float64); ok {
            shadowOffsetX = float32(__value_)
            hasShadowOffsetX = true
        }
    }
    var shadowOffsetY float32
    hasShadowOffsetY := false
    if __value, ok := jValue["shadowOffsetY"]; ok {
        if __value_, ok := __value.(float64); ok {
            shadowOffsetY = float32(__value_)
            hasShadowOffsetY = true
        }
    }
    var shadowBlurRadius int32
    hasShadowBlurRadius := false
    if __value, ok := jValue["shadowBlurRadius"]; ok {
        if __value_, ok := __value.(float64); ok {
            shadowBlurRadius = int32(__value_)
            hasShadowBlurRadius = true
        }
    }
    csbparser.TextOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.TextOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasFontResource {
        csbparser.TextOptionsAddFontResource(builder, fontResource)
    }
    if hasFontName {
        csbparser.TextOptionsAddFontName(builder, fontName)
    }
    if hasFontSize {
        csbparser.TextOptionsAddFontSize(builder, fontSize)
    }
    if hasText {
        csbparser.TextOptionsAddText(builder, text)
    }
    if hasAreaWidth {
        csbparser.TextOptionsAddAreaWidth(builder, areaWidth)
    }
    if hasAreaHeight {
        csbparser.TextOptionsAddAreaHeight(builder, areaHeight)
    }
    if hasHAlignment {
        csbparser.TextOptionsAddHAlignment(builder, hAlignment)
    }
    if hasVAlignment {
        csbparser.TextOptionsAddVAlignment(builder, vAlignment)
    }
    if hasTouchScaleEnable {
        csbparser.TextOptionsAddTouchScaleEnable(builder, touchScaleEnable)
    }
    if hasIsCustomSize {
        csbparser.TextOptionsAddIsCustomSize(builder, isCustomSize)
    }
    if hasOutlineEnabled {
        csbparser.TextOptionsAddOutlineEnabled(builder, outlineEnabled)
    }
    if hasOutlineColor {
        offset := CreateColor(builder, outlineColor)
        csbparser.TextOptionsAddOutlineColor(builder, offset)
    }
    if hasOutlineSize {
        csbparser.TextOptionsAddOutlineSize(builder, outlineSize)
    }
    if hasShadowEnabled {
        csbparser.TextOptionsAddShadowEnabled(builder, shadowEnabled)
    }
    if hasShadowColor {
        offset := CreateColor(builder, shadowColor)
        csbparser.TextOptionsAddShadowColor(builder, offset)
    }
    if hasShadowOffsetX {
        csbparser.TextOptionsAddShadowOffsetX(builder, shadowOffsetX)
    }
    if hasShadowOffsetY {
        csbparser.TextOptionsAddShadowOffsetY(builder, shadowOffsetY)
    }
    if hasShadowBlurRadius {
        csbparser.TextOptionsAddShadowBlurRadius(builder, shadowBlurRadius)
    }
    return csbparser.TextOptionsEnd(builder)
}
func CreateTextFieldOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var fontResource flatbuffers.UOffsetT
    hasFontResource := false
    if __value, ok := jValue["fontResource"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fontResource = CreateResourceData(builder, __value_)
            hasFontResource = true
        }
    }
    var fontName flatbuffers.UOffsetT
    hasFontName := false
    if __value, ok := jValue["fontName"]; ok {
        if __value_, ok := __value.(string); ok {
            fontName = builder.CreateString(__value_)
            hasFontName = true
        }
    }
    var fontSize int32
    hasFontSize := false
    if __value, ok := jValue["fontSize"]; ok {
        if __value_, ok := __value.(float64); ok {
            fontSize = int32(__value_)
            hasFontSize = true
        }
    }
    var text flatbuffers.UOffsetT
    hasText := false
    if __value, ok := jValue["text"]; ok {
        if __value_, ok := __value.(string); ok {
            text = builder.CreateString(__value_)
            hasText = true
        }
    }
    var placeHolder flatbuffers.UOffsetT
    hasPlaceHolder := false
    if __value, ok := jValue["placeHolder"]; ok {
        if __value_, ok := __value.(string); ok {
            placeHolder = builder.CreateString(__value_)
            hasPlaceHolder = true
        }
    }
    var passwordEnabled bool
    hasPasswordEnabled := false
    if __value, ok := jValue["passwordEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            passwordEnabled = bool(__value_)
            hasPasswordEnabled = true
        }
    }
    var passwordStyleText flatbuffers.UOffsetT
    hasPasswordStyleText := false
    if __value, ok := jValue["passwordStyleText"]; ok {
        if __value_, ok := __value.(string); ok {
            passwordStyleText = builder.CreateString(__value_)
            hasPasswordStyleText = true
        }
    }
    var maxLengthEnabled bool
    hasMaxLengthEnabled := false
    if __value, ok := jValue["maxLengthEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            maxLengthEnabled = bool(__value_)
            hasMaxLengthEnabled = true
        }
    }
    var maxLength int32
    hasMaxLength := false
    if __value, ok := jValue["maxLength"]; ok {
        if __value_, ok := __value.(float64); ok {
            maxLength = int32(__value_)
            hasMaxLength = true
        }
    }
    var areaWidth int32
    hasAreaWidth := false
    if __value, ok := jValue["areaWidth"]; ok {
        if __value_, ok := __value.(float64); ok {
            areaWidth = int32(__value_)
            hasAreaWidth = true
        }
    }
    var areaHeight int32
    hasAreaHeight := false
    if __value, ok := jValue["areaHeight"]; ok {
        if __value_, ok := __value.(float64); ok {
            areaHeight = int32(__value_)
            hasAreaHeight = true
        }
    }
    var isCustomSize bool
    hasIsCustomSize := false
    if __value, ok := jValue["isCustomSize"]; ok {
        if __value_, ok := __value.(bool); ok {
            isCustomSize = bool(__value_)
            hasIsCustomSize = true
        }
    }
    csbparser.TextFieldOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.TextFieldOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasFontResource {
        csbparser.TextFieldOptionsAddFontResource(builder, fontResource)
    }
    if hasFontName {
        csbparser.TextFieldOptionsAddFontName(builder, fontName)
    }
    if hasFontSize {
        csbparser.TextFieldOptionsAddFontSize(builder, fontSize)
    }
    if hasText {
        csbparser.TextFieldOptionsAddText(builder, text)
    }
    if hasPlaceHolder {
        csbparser.TextFieldOptionsAddPlaceHolder(builder, placeHolder)
    }
    if hasPasswordEnabled {
        csbparser.TextFieldOptionsAddPasswordEnabled(builder, passwordEnabled)
    }
    if hasPasswordStyleText {
        csbparser.TextFieldOptionsAddPasswordStyleText(builder, passwordStyleText)
    }
    if hasMaxLengthEnabled {
        csbparser.TextFieldOptionsAddMaxLengthEnabled(builder, maxLengthEnabled)
    }
    if hasMaxLength {
        csbparser.TextFieldOptionsAddMaxLength(builder, maxLength)
    }
    if hasAreaWidth {
        csbparser.TextFieldOptionsAddAreaWidth(builder, areaWidth)
    }
    if hasAreaHeight {
        csbparser.TextFieldOptionsAddAreaHeight(builder, areaHeight)
    }
    if hasIsCustomSize {
        csbparser.TextFieldOptionsAddIsCustomSize(builder, isCustomSize)
    }
    return csbparser.TextFieldOptionsEnd(builder)
}
func CreateLoadingBarOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var textureData flatbuffers.UOffsetT
    hasTextureData := false
    if __value, ok := jValue["textureData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            textureData = CreateResourceData(builder, __value_)
            hasTextureData = true
        }
    }
    var percent int32
    hasPercent := false
    if __value, ok := jValue["percent"]; ok {
        if __value_, ok := __value.(float64); ok {
            percent = int32(__value_)
            hasPercent = true
        }
    }
    var direction int32
    hasDirection := false
    if __value, ok := jValue["direction"]; ok {
        if __value_, ok := __value.(float64); ok {
            direction = int32(__value_)
            hasDirection = true
        }
    }
    csbparser.LoadingBarOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.LoadingBarOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasTextureData {
        csbparser.LoadingBarOptionsAddTextureData(builder, textureData)
    }
    if hasPercent {
        csbparser.LoadingBarOptionsAddPercent(builder, percent)
    }
    if hasDirection {
        csbparser.LoadingBarOptionsAddDirection(builder, direction)
    }
    return csbparser.LoadingBarOptionsEnd(builder)
}
func CreateSliderOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var barFileNameData flatbuffers.UOffsetT
    hasBarFileNameData := false
    if __value, ok := jValue["barFileNameData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            barFileNameData = CreateResourceData(builder, __value_)
            hasBarFileNameData = true
        }
    }
    var ballNormalData flatbuffers.UOffsetT
    hasBallNormalData := false
    if __value, ok := jValue["ballNormalData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            ballNormalData = CreateResourceData(builder, __value_)
            hasBallNormalData = true
        }
    }
    var ballPressedData flatbuffers.UOffsetT
    hasBallPressedData := false
    if __value, ok := jValue["ballPressedData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            ballPressedData = CreateResourceData(builder, __value_)
            hasBallPressedData = true
        }
    }
    var ballDisabledData flatbuffers.UOffsetT
    hasBallDisabledData := false
    if __value, ok := jValue["ballDisabledData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            ballDisabledData = CreateResourceData(builder, __value_)
            hasBallDisabledData = true
        }
    }
    var progressBarData flatbuffers.UOffsetT
    hasProgressBarData := false
    if __value, ok := jValue["progressBarData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            progressBarData = CreateResourceData(builder, __value_)
            hasProgressBarData = true
        }
    }
    var percent int32
    hasPercent := false
    if __value, ok := jValue["percent"]; ok {
        if __value_, ok := __value.(float64); ok {
            percent = int32(__value_)
            hasPercent = true
        }
    }
    var displaystate bool
    hasDisplaystate := false
    if __value, ok := jValue["displaystate"]; ok {
        if __value_, ok := __value.(bool); ok {
            displaystate = bool(__value_)
            hasDisplaystate = true
        }
    }
    csbparser.SliderOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.SliderOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasBarFileNameData {
        csbparser.SliderOptionsAddBarFileNameData(builder, barFileNameData)
    }
    if hasBallNormalData {
        csbparser.SliderOptionsAddBallNormalData(builder, ballNormalData)
    }
    if hasBallPressedData {
        csbparser.SliderOptionsAddBallPressedData(builder, ballPressedData)
    }
    if hasBallDisabledData {
        csbparser.SliderOptionsAddBallDisabledData(builder, ballDisabledData)
    }
    if hasProgressBarData {
        csbparser.SliderOptionsAddProgressBarData(builder, progressBarData)
    }
    if hasPercent {
        csbparser.SliderOptionsAddPercent(builder, percent)
    }
    if hasDisplaystate {
        csbparser.SliderOptionsAddDisplaystate(builder, displaystate)
    }
    return csbparser.SliderOptionsEnd(builder)
}
func CreatePanelOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var backGroundImageData flatbuffers.UOffsetT
    hasBackGroundImageData := false
    if __value, ok := jValue["backGroundImageData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            backGroundImageData = CreateResourceData(builder, __value_)
            hasBackGroundImageData = true
        }
    }
    var clipEnabled bool
    hasClipEnabled := false
    if __value, ok := jValue["clipEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            clipEnabled = bool(__value_)
            hasClipEnabled = true
        }
    }
    var bgColor map[string]interface{}
    hasBgColor := false
    if __value, ok := jValue["bgColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgColor = __value_
            hasBgColor = true
        }
    }
    var bgStartColor map[string]interface{}
    hasBgStartColor := false
    if __value, ok := jValue["bgStartColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgStartColor = __value_
            hasBgStartColor = true
        }
    }
    var bgEndColor map[string]interface{}
    hasBgEndColor := false
    if __value, ok := jValue["bgEndColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgEndColor = __value_
            hasBgEndColor = true
        }
    }
    var colorType int32
    hasColorType := false
    if __value, ok := jValue["colorType"]; ok {
        if __value_, ok := __value.(float64); ok {
            colorType = int32(__value_)
            hasColorType = true
        }
    }
    var bgColorOpacity byte
    hasBgColorOpacity := false
    if __value, ok := jValue["bgColorOpacity"]; ok {
        if __value_, ok := __value.(float64); ok {
            bgColorOpacity = byte(__value_)
            hasBgColorOpacity = true
        }
    }
    var colorVector map[string]interface{}
    hasColorVector := false
    if __value, ok := jValue["colorVector"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            colorVector = __value_
            hasColorVector = true
        }
    }
    var capInsets map[string]interface{}
    hasCapInsets := false
    if __value, ok := jValue["capInsets"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            capInsets = __value_
            hasCapInsets = true
        }
    }
    var scale9Size map[string]interface{}
    hasScale9Size := false
    if __value, ok := jValue["scale9Size"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scale9Size = __value_
            hasScale9Size = true
        }
    }
    var backGroundScale9Enabled bool
    hasBackGroundScale9Enabled := false
    if __value, ok := jValue["backGroundScale9Enabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            backGroundScale9Enabled = bool(__value_)
            hasBackGroundScale9Enabled = true
        }
    }
    csbparser.PanelOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.PanelOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasBackGroundImageData {
        csbparser.PanelOptionsAddBackGroundImageData(builder, backGroundImageData)
    }
    if hasClipEnabled {
        csbparser.PanelOptionsAddClipEnabled(builder, clipEnabled)
    }
    if hasBgColor {
        offset := CreateColor(builder, bgColor)
        csbparser.PanelOptionsAddBgColor(builder, offset)
    }
    if hasBgStartColor {
        offset := CreateColor(builder, bgStartColor)
        csbparser.PanelOptionsAddBgStartColor(builder, offset)
    }
    if hasBgEndColor {
        offset := CreateColor(builder, bgEndColor)
        csbparser.PanelOptionsAddBgEndColor(builder, offset)
    }
    if hasColorType {
        csbparser.PanelOptionsAddColorType(builder, colorType)
    }
    if hasBgColorOpacity {
        csbparser.PanelOptionsAddBgColorOpacity(builder, bgColorOpacity)
    }
    if hasColorVector {
        offset := CreateColorVector(builder, colorVector)
        csbparser.PanelOptionsAddColorVector(builder, offset)
    }
    if hasCapInsets {
        offset := CreateCapInsets(builder, capInsets)
        csbparser.PanelOptionsAddCapInsets(builder, offset)
    }
    if hasScale9Size {
        offset := CreateFlatSize(builder, scale9Size)
        csbparser.PanelOptionsAddScale9Size(builder, offset)
    }
    if hasBackGroundScale9Enabled {
        csbparser.PanelOptionsAddBackGroundScale9Enabled(builder, backGroundScale9Enabled)
    }
    return csbparser.PanelOptionsEnd(builder)
}
func CreateScrollViewOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var backGroundImageData flatbuffers.UOffsetT
    hasBackGroundImageData := false
    if __value, ok := jValue["backGroundImageData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            backGroundImageData = CreateResourceData(builder, __value_)
            hasBackGroundImageData = true
        }
    }
    var clipEnabled bool
    hasClipEnabled := false
    if __value, ok := jValue["clipEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            clipEnabled = bool(__value_)
            hasClipEnabled = true
        }
    }
    var bgColor map[string]interface{}
    hasBgColor := false
    if __value, ok := jValue["bgColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgColor = __value_
            hasBgColor = true
        }
    }
    var bgStartColor map[string]interface{}
    hasBgStartColor := false
    if __value, ok := jValue["bgStartColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgStartColor = __value_
            hasBgStartColor = true
        }
    }
    var bgEndColor map[string]interface{}
    hasBgEndColor := false
    if __value, ok := jValue["bgEndColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgEndColor = __value_
            hasBgEndColor = true
        }
    }
    var colorType int32
    hasColorType := false
    if __value, ok := jValue["colorType"]; ok {
        if __value_, ok := __value.(float64); ok {
            colorType = int32(__value_)
            hasColorType = true
        }
    }
    var bgColorOpacity byte
    hasBgColorOpacity := false
    if __value, ok := jValue["bgColorOpacity"]; ok {
        if __value_, ok := __value.(float64); ok {
            bgColorOpacity = byte(__value_)
            hasBgColorOpacity = true
        }
    }
    var colorVector map[string]interface{}
    hasColorVector := false
    if __value, ok := jValue["colorVector"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            colorVector = __value_
            hasColorVector = true
        }
    }
    var capInsets map[string]interface{}
    hasCapInsets := false
    if __value, ok := jValue["capInsets"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            capInsets = __value_
            hasCapInsets = true
        }
    }
    var scale9Size map[string]interface{}
    hasScale9Size := false
    if __value, ok := jValue["scale9Size"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scale9Size = __value_
            hasScale9Size = true
        }
    }
    var backGroundScale9Enabled bool
    hasBackGroundScale9Enabled := false
    if __value, ok := jValue["backGroundScale9Enabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            backGroundScale9Enabled = bool(__value_)
            hasBackGroundScale9Enabled = true
        }
    }
    var innerSize map[string]interface{}
    hasInnerSize := false
    if __value, ok := jValue["innerSize"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            innerSize = __value_
            hasInnerSize = true
        }
    }
    var direction int32
    hasDirection := false
    if __value, ok := jValue["direction"]; ok {
        if __value_, ok := __value.(float64); ok {
            direction = int32(__value_)
            hasDirection = true
        }
    }
    var bounceEnabled bool
    hasBounceEnabled := false
    if __value, ok := jValue["bounceEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            bounceEnabled = bool(__value_)
            hasBounceEnabled = true
        }
    }
    var scrollbarEnabeld bool
    hasScrollbarEnabeld := false
    if __value, ok := jValue["scrollbarEnabeld"]; ok {
        if __value_, ok := __value.(bool); ok {
            scrollbarEnabeld = bool(__value_)
            hasScrollbarEnabeld = true
        }
    }
    var scrollbarAutoHide bool
    hasScrollbarAutoHide := false
    if __value, ok := jValue["scrollbarAutoHide"]; ok {
        if __value_, ok := __value.(bool); ok {
            scrollbarAutoHide = bool(__value_)
            hasScrollbarAutoHide = true
        }
    }
    var scrollbarAutoHideTime float32
    hasScrollbarAutoHideTime := false
    if __value, ok := jValue["scrollbarAutoHideTime"]; ok {
        if __value_, ok := __value.(float64); ok {
            scrollbarAutoHideTime = float32(__value_)
            hasScrollbarAutoHideTime = true
        }
    }
    csbparser.ScrollViewOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.ScrollViewOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasBackGroundImageData {
        csbparser.ScrollViewOptionsAddBackGroundImageData(builder, backGroundImageData)
    }
    if hasClipEnabled {
        csbparser.ScrollViewOptionsAddClipEnabled(builder, clipEnabled)
    }
    if hasBgColor {
        offset := CreateColor(builder, bgColor)
        csbparser.ScrollViewOptionsAddBgColor(builder, offset)
    }
    if hasBgStartColor {
        offset := CreateColor(builder, bgStartColor)
        csbparser.ScrollViewOptionsAddBgStartColor(builder, offset)
    }
    if hasBgEndColor {
        offset := CreateColor(builder, bgEndColor)
        csbparser.ScrollViewOptionsAddBgEndColor(builder, offset)
    }
    if hasColorType {
        csbparser.ScrollViewOptionsAddColorType(builder, colorType)
    }
    if hasBgColorOpacity {
        csbparser.ScrollViewOptionsAddBgColorOpacity(builder, bgColorOpacity)
    }
    if hasColorVector {
        offset := CreateColorVector(builder, colorVector)
        csbparser.ScrollViewOptionsAddColorVector(builder, offset)
    }
    if hasCapInsets {
        offset := CreateCapInsets(builder, capInsets)
        csbparser.ScrollViewOptionsAddCapInsets(builder, offset)
    }
    if hasScale9Size {
        offset := CreateFlatSize(builder, scale9Size)
        csbparser.ScrollViewOptionsAddScale9Size(builder, offset)
    }
    if hasBackGroundScale9Enabled {
        csbparser.ScrollViewOptionsAddBackGroundScale9Enabled(builder, backGroundScale9Enabled)
    }
    if hasInnerSize {
        offset := CreateFlatSize(builder, innerSize)
        csbparser.ScrollViewOptionsAddInnerSize(builder, offset)
    }
    if hasDirection {
        csbparser.ScrollViewOptionsAddDirection(builder, direction)
    }
    if hasBounceEnabled {
        csbparser.ScrollViewOptionsAddBounceEnabled(builder, bounceEnabled)
    }
    if hasScrollbarEnabeld {
        csbparser.ScrollViewOptionsAddScrollbarEnabeld(builder, scrollbarEnabeld)
    }
    if hasScrollbarAutoHide {
        csbparser.ScrollViewOptionsAddScrollbarAutoHide(builder, scrollbarAutoHide)
    }
    if hasScrollbarAutoHideTime {
        csbparser.ScrollViewOptionsAddScrollbarAutoHideTime(builder, scrollbarAutoHideTime)
    }
    return csbparser.ScrollViewOptionsEnd(builder)
}
func CreatePageViewOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var backGroundImageData flatbuffers.UOffsetT
    hasBackGroundImageData := false
    if __value, ok := jValue["backGroundImageData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            backGroundImageData = CreateResourceData(builder, __value_)
            hasBackGroundImageData = true
        }
    }
    var clipEnabled bool
    hasClipEnabled := false
    if __value, ok := jValue["clipEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            clipEnabled = bool(__value_)
            hasClipEnabled = true
        }
    }
    var bgColor map[string]interface{}
    hasBgColor := false
    if __value, ok := jValue["bgColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgColor = __value_
            hasBgColor = true
        }
    }
    var bgStartColor map[string]interface{}
    hasBgStartColor := false
    if __value, ok := jValue["bgStartColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgStartColor = __value_
            hasBgStartColor = true
        }
    }
    var bgEndColor map[string]interface{}
    hasBgEndColor := false
    if __value, ok := jValue["bgEndColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgEndColor = __value_
            hasBgEndColor = true
        }
    }
    var colorType int32
    hasColorType := false
    if __value, ok := jValue["colorType"]; ok {
        if __value_, ok := __value.(float64); ok {
            colorType = int32(__value_)
            hasColorType = true
        }
    }
    var bgColorOpacity byte
    hasBgColorOpacity := false
    if __value, ok := jValue["bgColorOpacity"]; ok {
        if __value_, ok := __value.(float64); ok {
            bgColorOpacity = byte(__value_)
            hasBgColorOpacity = true
        }
    }
    var colorVector map[string]interface{}
    hasColorVector := false
    if __value, ok := jValue["colorVector"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            colorVector = __value_
            hasColorVector = true
        }
    }
    var capInsets map[string]interface{}
    hasCapInsets := false
    if __value, ok := jValue["capInsets"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            capInsets = __value_
            hasCapInsets = true
        }
    }
    var scale9Size map[string]interface{}
    hasScale9Size := false
    if __value, ok := jValue["scale9Size"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scale9Size = __value_
            hasScale9Size = true
        }
    }
    var backGroundScale9Enabled bool
    hasBackGroundScale9Enabled := false
    if __value, ok := jValue["backGroundScale9Enabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            backGroundScale9Enabled = bool(__value_)
            hasBackGroundScale9Enabled = true
        }
    }
    csbparser.PageViewOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.PageViewOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasBackGroundImageData {
        csbparser.PageViewOptionsAddBackGroundImageData(builder, backGroundImageData)
    }
    if hasClipEnabled {
        csbparser.PageViewOptionsAddClipEnabled(builder, clipEnabled)
    }
    if hasBgColor {
        offset := CreateColor(builder, bgColor)
        csbparser.PageViewOptionsAddBgColor(builder, offset)
    }
    if hasBgStartColor {
        offset := CreateColor(builder, bgStartColor)
        csbparser.PageViewOptionsAddBgStartColor(builder, offset)
    }
    if hasBgEndColor {
        offset := CreateColor(builder, bgEndColor)
        csbparser.PageViewOptionsAddBgEndColor(builder, offset)
    }
    if hasColorType {
        csbparser.PageViewOptionsAddColorType(builder, colorType)
    }
    if hasBgColorOpacity {
        csbparser.PageViewOptionsAddBgColorOpacity(builder, bgColorOpacity)
    }
    if hasColorVector {
        offset := CreateColorVector(builder, colorVector)
        csbparser.PageViewOptionsAddColorVector(builder, offset)
    }
    if hasCapInsets {
        offset := CreateCapInsets(builder, capInsets)
        csbparser.PageViewOptionsAddCapInsets(builder, offset)
    }
    if hasScale9Size {
        offset := CreateFlatSize(builder, scale9Size)
        csbparser.PageViewOptionsAddScale9Size(builder, offset)
    }
    if hasBackGroundScale9Enabled {
        csbparser.PageViewOptionsAddBackGroundScale9Enabled(builder, backGroundScale9Enabled)
    }
    return csbparser.PageViewOptionsEnd(builder)
}
func CreateListViewOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var widgetOptions flatbuffers.UOffsetT
    hasWidgetOptions := false
    if __value, ok := jValue["widgetOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            widgetOptions = CreateWidgetOptions(builder, __value_)
            hasWidgetOptions = true
        }
    }
    var backGroundImageData flatbuffers.UOffsetT
    hasBackGroundImageData := false
    if __value, ok := jValue["backGroundImageData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            backGroundImageData = CreateResourceData(builder, __value_)
            hasBackGroundImageData = true
        }
    }
    var clipEnabled bool
    hasClipEnabled := false
    if __value, ok := jValue["clipEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            clipEnabled = bool(__value_)
            hasClipEnabled = true
        }
    }
    var bgColor map[string]interface{}
    hasBgColor := false
    if __value, ok := jValue["bgColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgColor = __value_
            hasBgColor = true
        }
    }
    var bgStartColor map[string]interface{}
    hasBgStartColor := false
    if __value, ok := jValue["bgStartColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgStartColor = __value_
            hasBgStartColor = true
        }
    }
    var bgEndColor map[string]interface{}
    hasBgEndColor := false
    if __value, ok := jValue["bgEndColor"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            bgEndColor = __value_
            hasBgEndColor = true
        }
    }
    var colorType int32
    hasColorType := false
    if __value, ok := jValue["colorType"]; ok {
        if __value_, ok := __value.(float64); ok {
            colorType = int32(__value_)
            hasColorType = true
        }
    }
    var bgColorOpacity byte
    hasBgColorOpacity := false
    if __value, ok := jValue["bgColorOpacity"]; ok {
        if __value_, ok := __value.(float64); ok {
            bgColorOpacity = byte(__value_)
            hasBgColorOpacity = true
        }
    }
    var colorVector map[string]interface{}
    hasColorVector := false
    if __value, ok := jValue["colorVector"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            colorVector = __value_
            hasColorVector = true
        }
    }
    var capInsets map[string]interface{}
    hasCapInsets := false
    if __value, ok := jValue["capInsets"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            capInsets = __value_
            hasCapInsets = true
        }
    }
    var scale9Size map[string]interface{}
    hasScale9Size := false
    if __value, ok := jValue["scale9Size"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scale9Size = __value_
            hasScale9Size = true
        }
    }
    var backGroundScale9Enabled bool
    hasBackGroundScale9Enabled := false
    if __value, ok := jValue["backGroundScale9Enabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            backGroundScale9Enabled = bool(__value_)
            hasBackGroundScale9Enabled = true
        }
    }
    var innerSize map[string]interface{}
    hasInnerSize := false
    if __value, ok := jValue["innerSize"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            innerSize = __value_
            hasInnerSize = true
        }
    }
    var direction int32
    hasDirection := false
    if __value, ok := jValue["direction"]; ok {
        if __value_, ok := __value.(float64); ok {
            direction = int32(__value_)
            hasDirection = true
        }
    }
    var bounceEnabled bool
    hasBounceEnabled := false
    if __value, ok := jValue["bounceEnabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            bounceEnabled = bool(__value_)
            hasBounceEnabled = true
        }
    }
    var itemMargin int32
    hasItemMargin := false
    if __value, ok := jValue["itemMargin"]; ok {
        if __value_, ok := __value.(float64); ok {
            itemMargin = int32(__value_)
            hasItemMargin = true
        }
    }
    var directionType flatbuffers.UOffsetT
    hasDirectionType := false
    if __value, ok := jValue["directionType"]; ok {
        if __value_, ok := __value.(string); ok {
            directionType = builder.CreateString(__value_)
            hasDirectionType = true
        }
    }
    var horizontalType flatbuffers.UOffsetT
    hasHorizontalType := false
    if __value, ok := jValue["horizontalType"]; ok {
        if __value_, ok := __value.(string); ok {
            horizontalType = builder.CreateString(__value_)
            hasHorizontalType = true
        }
    }
    var verticalType flatbuffers.UOffsetT
    hasVerticalType := false
    if __value, ok := jValue["verticalType"]; ok {
        if __value_, ok := __value.(string); ok {
            verticalType = builder.CreateString(__value_)
            hasVerticalType = true
        }
    }
    csbparser.ListViewOptionsStart(builder)
    if hasWidgetOptions {
        csbparser.ListViewOptionsAddWidgetOptions(builder, widgetOptions)
    }
    if hasBackGroundImageData {
        csbparser.ListViewOptionsAddBackGroundImageData(builder, backGroundImageData)
    }
    if hasClipEnabled {
        csbparser.ListViewOptionsAddClipEnabled(builder, clipEnabled)
    }
    if hasBgColor {
        offset := CreateColor(builder, bgColor)
        csbparser.ListViewOptionsAddBgColor(builder, offset)
    }
    if hasBgStartColor {
        offset := CreateColor(builder, bgStartColor)
        csbparser.ListViewOptionsAddBgStartColor(builder, offset)
    }
    if hasBgEndColor {
        offset := CreateColor(builder, bgEndColor)
        csbparser.ListViewOptionsAddBgEndColor(builder, offset)
    }
    if hasColorType {
        csbparser.ListViewOptionsAddColorType(builder, colorType)
    }
    if hasBgColorOpacity {
        csbparser.ListViewOptionsAddBgColorOpacity(builder, bgColorOpacity)
    }
    if hasColorVector {
        offset := CreateColorVector(builder, colorVector)
        csbparser.ListViewOptionsAddColorVector(builder, offset)
    }
    if hasCapInsets {
        offset := CreateCapInsets(builder, capInsets)
        csbparser.ListViewOptionsAddCapInsets(builder, offset)
    }
    if hasScale9Size {
        offset := CreateFlatSize(builder, scale9Size)
        csbparser.ListViewOptionsAddScale9Size(builder, offset)
    }
    if hasBackGroundScale9Enabled {
        csbparser.ListViewOptionsAddBackGroundScale9Enabled(builder, backGroundScale9Enabled)
    }
    if hasInnerSize {
        offset := CreateFlatSize(builder, innerSize)
        csbparser.ListViewOptionsAddInnerSize(builder, offset)
    }
    if hasDirection {
        csbparser.ListViewOptionsAddDirection(builder, direction)
    }
    if hasBounceEnabled {
        csbparser.ListViewOptionsAddBounceEnabled(builder, bounceEnabled)
    }
    if hasItemMargin {
        csbparser.ListViewOptionsAddItemMargin(builder, itemMargin)
    }
    if hasDirectionType {
        csbparser.ListViewOptionsAddDirectionType(builder, directionType)
    }
    if hasHorizontalType {
        csbparser.ListViewOptionsAddHorizontalType(builder, horizontalType)
    }
    if hasVerticalType {
        csbparser.ListViewOptionsAddVerticalType(builder, verticalType)
    }
    return csbparser.ListViewOptionsEnd(builder)
}
func CreateProjectNodeOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var nodeOptions flatbuffers.UOffsetT
    hasNodeOptions := false
    if __value, ok := jValue["nodeOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            nodeOptions = CreateWidgetOptions(builder, __value_)
            hasNodeOptions = true
        }
    }
    var fileName flatbuffers.UOffsetT
    hasFileName := false
    if __value, ok := jValue["fileName"]; ok {
        if __value_, ok := __value.(string); ok {
            fileName = builder.CreateString(__value_)
            hasFileName = true
        }
    }
    var innerActionSpeed float32
    hasInnerActionSpeed := false
    if __value, ok := jValue["innerActionSpeed"]; ok {
        if __value_, ok := __value.(float64); ok {
            innerActionSpeed = float32(__value_)
            hasInnerActionSpeed = true
        }
    }
    csbparser.ProjectNodeOptionsStart(builder)
    if hasNodeOptions {
        csbparser.ProjectNodeOptionsAddNodeOptions(builder, nodeOptions)
    }
    if hasFileName {
        csbparser.ProjectNodeOptionsAddFileName(builder, fileName)
    }
    if hasInnerActionSpeed {
        csbparser.ProjectNodeOptionsAddInnerActionSpeed(builder, innerActionSpeed)
    }
    return csbparser.ProjectNodeOptionsEnd(builder)
}
func CreateComponentOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var nodeOptions flatbuffers.UOffsetT
    hasNodeOptions := false
    if __value, ok := jValue["nodeOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            nodeOptions = CreateWidgetOptions(builder, __value_)
            hasNodeOptions = true
        }
    }
    var comType flatbuffers.UOffsetT
    hasComType := false
    if __value, ok := jValue["comType"]; ok {
        if __value_, ok := __value.(string); ok {
            comType = builder.CreateString(__value_)
            hasComType = true
        }
    }
    var comAudioOptions flatbuffers.UOffsetT
    hasComAudioOptions := false
    if __value, ok := jValue["comAudioOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            comAudioOptions = CreateComAudioOptions(builder, __value_)
            hasComAudioOptions = true
        }
    }
    csbparser.ComponentOptionsStart(builder)
    if hasNodeOptions {
        csbparser.ComponentOptionsAddNodeOptions(builder, nodeOptions)
    }
    if hasComType {
        csbparser.ComponentOptionsAddComType(builder, comType)
    }
    if hasComAudioOptions {
        csbparser.ComponentOptionsAddComAudioOptions(builder, comAudioOptions)
    }
    return csbparser.ComponentOptionsEnd(builder)
}
func CreateComAudioOptions(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var nodeOptions flatbuffers.UOffsetT
    hasNodeOptions := false
    if __value, ok := jValue["nodeOptions"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            nodeOptions = CreateWidgetOptions(builder, __value_)
            hasNodeOptions = true
        }
    }
    var name flatbuffers.UOffsetT
    hasName := false
    if __value, ok := jValue["name"]; ok {
        if __value_, ok := __value.(string); ok {
            name = builder.CreateString(__value_)
            hasName = true
        }
    }
    var enabled bool
    hasEnabled := false
    if __value, ok := jValue["enabled"]; ok {
        if __value_, ok := __value.(bool); ok {
            enabled = bool(__value_)
            hasEnabled = true
        }
    }
    var loop bool
    hasLoop := false
    if __value, ok := jValue["loop"]; ok {
        if __value_, ok := __value.(bool); ok {
            loop = bool(__value_)
            hasLoop = true
        }
    }
    var volume int32
    hasVolume := false
    if __value, ok := jValue["volume"]; ok {
        if __value_, ok := __value.(float64); ok {
            volume = int32(__value_)
            hasVolume = true
        }
    }
    var fileNameData flatbuffers.UOffsetT
    hasFileNameData := false
    if __value, ok := jValue["fileNameData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            fileNameData = CreateResourceData(builder, __value_)
            hasFileNameData = true
        }
    }
    csbparser.ComAudioOptionsStart(builder)
    if hasNodeOptions {
        csbparser.ComAudioOptionsAddNodeOptions(builder, nodeOptions)
    }
    if hasName {
        csbparser.ComAudioOptionsAddName(builder, name)
    }
    if hasEnabled {
        csbparser.ComAudioOptionsAddEnabled(builder, enabled)
    }
    if hasLoop {
        csbparser.ComAudioOptionsAddLoop(builder, loop)
    }
    if hasVolume {
        csbparser.ComAudioOptionsAddVolume(builder, volume)
    }
    if hasFileNameData {
        csbparser.ComAudioOptionsAddFileNameData(builder, fileNameData)
    }
    return csbparser.ComAudioOptionsEnd(builder)
}
func CreateAnimationInfo(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var name flatbuffers.UOffsetT
    hasName := false
    if __value, ok := jValue["name"]; ok {
        if __value_, ok := __value.(string); ok {
            name = builder.CreateString(__value_)
            hasName = true
        }
    }
    var startIndex int32
    hasStartIndex := false
    if __value, ok := jValue["startIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            startIndex = int32(__value_)
            hasStartIndex = true
        }
    }
    var endIndex int32
    hasEndIndex := false
    if __value, ok := jValue["endIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            endIndex = int32(__value_)
            hasEndIndex = true
        }
    }
    csbparser.AnimationInfoStart(builder)
    if hasName {
        csbparser.AnimationInfoAddName(builder, name)
    }
    if hasStartIndex {
        csbparser.AnimationInfoAddStartIndex(builder, startIndex)
    }
    if hasEndIndex {
        csbparser.AnimationInfoAddEndIndex(builder, endIndex)
    }
    return csbparser.AnimationInfoEnd(builder)
}
func CreateNodeAction(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var duration int32
    hasDuration := false
    if __value, ok := jValue["duration"]; ok {
        if __value_, ok := __value.(float64); ok {
            duration = int32(__value_)
            hasDuration = true
        }
    }
    var speed float32
    hasSpeed := false
    if __value, ok := jValue["speed"]; ok {
        if __value_, ok := __value.(float64); ok {
            speed = float32(__value_)
            hasSpeed = true
        }
    }
    var timeLines flatbuffers.UOffsetT
    hasTimeLines := false
    if __value, ok := jValue["timeLines"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(map[string]interface{}); ok {
                    vector = append(vector, CreateTimeLine(builder, v_))
                }
            }
            csbparser.NodeActionStartTimeLinesVector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            timeLines = builder.EndVector(len(vector))
            hasTimeLines = true
        }
    }
    var currentAnimationName flatbuffers.UOffsetT
    hasCurrentAnimationName := false
    if __value, ok := jValue["currentAnimationName"]; ok {
        if __value_, ok := __value.(string); ok {
            currentAnimationName = builder.CreateString(__value_)
            hasCurrentAnimationName = true
        }
    }
    csbparser.NodeActionStart(builder)
    if hasDuration {
        csbparser.NodeActionAddDuration(builder, duration)
    }
    if hasSpeed {
        csbparser.NodeActionAddSpeed(builder, speed)
    }
    if hasTimeLines {
        csbparser.NodeActionAddTimeLines(builder, timeLines)
    }
    if hasCurrentAnimationName {
        csbparser.NodeActionAddCurrentAnimationName(builder, currentAnimationName)
    }
    return csbparser.NodeActionEnd(builder)
}
func CreateTimeLine(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var property flatbuffers.UOffsetT
    hasProperty := false
    if __value, ok := jValue["property"]; ok {
        if __value_, ok := __value.(string); ok {
            property = builder.CreateString(__value_)
            hasProperty = true
        }
    }
    var actionTag int32
    hasActionTag := false
    if __value, ok := jValue["actionTag"]; ok {
        if __value_, ok := __value.(float64); ok {
            actionTag = int32(__value_)
            hasActionTag = true
        }
    }
    var frames flatbuffers.UOffsetT
    hasFrames := false
    if __value, ok := jValue["frames"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(map[string]interface{}); ok {
                    vector = append(vector, CreateFrame(builder, v_))
                }
            }
            csbparser.TimeLineStartFramesVector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            frames = builder.EndVector(len(vector))
            hasFrames = true
        }
    }
    csbparser.TimeLineStart(builder)
    if hasProperty {
        csbparser.TimeLineAddProperty(builder, property)
    }
    if hasActionTag {
        csbparser.TimeLineAddActionTag(builder, actionTag)
    }
    if hasFrames {
        csbparser.TimeLineAddFrames(builder, frames)
    }
    return csbparser.TimeLineEnd(builder)
}
func CreateFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var pointFrame flatbuffers.UOffsetT
    hasPointFrame := false
    if __value, ok := jValue["pointFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            pointFrame = CreatePointFrame(builder, __value_)
            hasPointFrame = true
        }
    }
    var scaleFrame flatbuffers.UOffsetT
    hasScaleFrame := false
    if __value, ok := jValue["scaleFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scaleFrame = CreateScaleFrame(builder, __value_)
            hasScaleFrame = true
        }
    }
    var colorFrame flatbuffers.UOffsetT
    hasColorFrame := false
    if __value, ok := jValue["colorFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            colorFrame = CreateColorFrame(builder, __value_)
            hasColorFrame = true
        }
    }
    var textureFrame flatbuffers.UOffsetT
    hasTextureFrame := false
    if __value, ok := jValue["textureFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            textureFrame = CreateTextureFrame(builder, __value_)
            hasTextureFrame = true
        }
    }
    var eventFrame flatbuffers.UOffsetT
    hasEventFrame := false
    if __value, ok := jValue["eventFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            eventFrame = CreateEventFrame(builder, __value_)
            hasEventFrame = true
        }
    }
    var intFrame flatbuffers.UOffsetT
    hasIntFrame := false
    if __value, ok := jValue["intFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            intFrame = CreateIntFrame(builder, __value_)
            hasIntFrame = true
        }
    }
    var boolFrame flatbuffers.UOffsetT
    hasBoolFrame := false
    if __value, ok := jValue["boolFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            boolFrame = CreateBoolFrame(builder, __value_)
            hasBoolFrame = true
        }
    }
    var innerActionFrame flatbuffers.UOffsetT
    hasInnerActionFrame := false
    if __value, ok := jValue["innerActionFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            innerActionFrame = CreateInnerActionFrame(builder, __value_)
            hasInnerActionFrame = true
        }
    }
    var blendFrame flatbuffers.UOffsetT
    hasBlendFrame := false
    if __value, ok := jValue["blendFrame"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            blendFrame = CreateBlendFrame(builder, __value_)
            hasBlendFrame = true
        }
    }
    csbparser.FrameStart(builder)
    if hasPointFrame {
        csbparser.FrameAddPointFrame(builder, pointFrame)
    }
    if hasScaleFrame {
        csbparser.FrameAddScaleFrame(builder, scaleFrame)
    }
    if hasColorFrame {
        csbparser.FrameAddColorFrame(builder, colorFrame)
    }
    if hasTextureFrame {
        csbparser.FrameAddTextureFrame(builder, textureFrame)
    }
    if hasEventFrame {
        csbparser.FrameAddEventFrame(builder, eventFrame)
    }
    if hasIntFrame {
        csbparser.FrameAddIntFrame(builder, intFrame)
    }
    if hasBoolFrame {
        csbparser.FrameAddBoolFrame(builder, boolFrame)
    }
    if hasInnerActionFrame {
        csbparser.FrameAddInnerActionFrame(builder, innerActionFrame)
    }
    if hasBlendFrame {
        csbparser.FrameAddBlendFrame(builder, blendFrame)
    }
    return csbparser.FrameEnd(builder)
}
func CreatePointFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var position map[string]interface{}
    hasPosition := false
    if __value, ok := jValue["position"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            position = __value_
            hasPosition = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.PointFrameStart(builder)
    if hasFrameIndex {
        csbparser.PointFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.PointFrameAddTween(builder, tween)
    }
    if hasPosition {
        offset := CreatePosition(builder, position)
        csbparser.PointFrameAddPosition(builder, offset)
    }
    if hasEasingData {
        csbparser.PointFrameAddEasingData(builder, easingData)
    }
    return csbparser.PointFrameEnd(builder)
}
func CreateScaleFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var scale map[string]interface{}
    hasScale := false
    if __value, ok := jValue["scale"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            scale = __value_
            hasScale = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.ScaleFrameStart(builder)
    if hasFrameIndex {
        csbparser.ScaleFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.ScaleFrameAddTween(builder, tween)
    }
    if hasScale {
        offset := CreateScale(builder, scale)
        csbparser.ScaleFrameAddScale(builder, offset)
    }
    if hasEasingData {
        csbparser.ScaleFrameAddEasingData(builder, easingData)
    }
    return csbparser.ScaleFrameEnd(builder)
}
func CreateColorFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var color map[string]interface{}
    hasColor := false
    if __value, ok := jValue["color"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            color = __value_
            hasColor = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.ColorFrameStart(builder)
    if hasFrameIndex {
        csbparser.ColorFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.ColorFrameAddTween(builder, tween)
    }
    if hasColor {
        offset := CreateColor(builder, color)
        csbparser.ColorFrameAddColor(builder, offset)
    }
    if hasEasingData {
        csbparser.ColorFrameAddEasingData(builder, easingData)
    }
    return csbparser.ColorFrameEnd(builder)
}
func CreateTextureFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var textureFile flatbuffers.UOffsetT
    hasTextureFile := false
    if __value, ok := jValue["textureFile"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            textureFile = CreateResourceData(builder, __value_)
            hasTextureFile = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.TextureFrameStart(builder)
    if hasFrameIndex {
        csbparser.TextureFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.TextureFrameAddTween(builder, tween)
    }
    if hasTextureFile {
        csbparser.TextureFrameAddTextureFile(builder, textureFile)
    }
    if hasEasingData {
        csbparser.TextureFrameAddEasingData(builder, easingData)
    }
    return csbparser.TextureFrameEnd(builder)
}
func CreateEventFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var value flatbuffers.UOffsetT
    hasValue := false
    if __value, ok := jValue["value"]; ok {
        if __value_, ok := __value.(string); ok {
            value = builder.CreateString(__value_)
            hasValue = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.EventFrameStart(builder)
    if hasFrameIndex {
        csbparser.EventFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.EventFrameAddTween(builder, tween)
    }
    if hasValue {
        csbparser.EventFrameAddValue(builder, value)
    }
    if hasEasingData {
        csbparser.EventFrameAddEasingData(builder, easingData)
    }
    return csbparser.EventFrameEnd(builder)
}
func CreateIntFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var value int32
    hasValue := false
    if __value, ok := jValue["value"]; ok {
        if __value_, ok := __value.(float64); ok {
            value = int32(__value_)
            hasValue = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.IntFrameStart(builder)
    if hasFrameIndex {
        csbparser.IntFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.IntFrameAddTween(builder, tween)
    }
    if hasValue {
        csbparser.IntFrameAddValue(builder, value)
    }
    if hasEasingData {
        csbparser.IntFrameAddEasingData(builder, easingData)
    }
    return csbparser.IntFrameEnd(builder)
}
func CreateBoolFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var value bool
    hasValue := false
    if __value, ok := jValue["value"]; ok {
        if __value_, ok := __value.(bool); ok {
            value = bool(__value_)
            hasValue = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.BoolFrameStart(builder)
    if hasFrameIndex {
        csbparser.BoolFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.BoolFrameAddTween(builder, tween)
    }
    if hasValue {
        csbparser.BoolFrameAddValue(builder, value)
    }
    if hasEasingData {
        csbparser.BoolFrameAddEasingData(builder, easingData)
    }
    return csbparser.BoolFrameEnd(builder)
}
func CreateInnerActionFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var innerActionType int32
    hasInnerActionType := false
    if __value, ok := jValue["innerActionType"]; ok {
        if __value_, ok := __value.(float64); ok {
            innerActionType = int32(__value_)
            hasInnerActionType = true
        }
    }
    var currentAniamtionName flatbuffers.UOffsetT
    hasCurrentAniamtionName := false
    if __value, ok := jValue["currentAniamtionName"]; ok {
        if __value_, ok := __value.(string); ok {
            currentAniamtionName = builder.CreateString(__value_)
            hasCurrentAniamtionName = true
        }
    }
    var singleFrameIndex int32
    hasSingleFrameIndex := false
    if __value, ok := jValue["singleFrameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            singleFrameIndex = int32(__value_)
            hasSingleFrameIndex = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.InnerActionFrameStart(builder)
    if hasFrameIndex {
        csbparser.InnerActionFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.InnerActionFrameAddTween(builder, tween)
    }
    if hasInnerActionType {
        csbparser.InnerActionFrameAddInnerActionType(builder, innerActionType)
    }
    if hasCurrentAniamtionName {
        csbparser.InnerActionFrameAddCurrentAniamtionName(builder, currentAniamtionName)
    }
    if hasSingleFrameIndex {
        csbparser.InnerActionFrameAddSingleFrameIndex(builder, singleFrameIndex)
    }
    if hasEasingData {
        csbparser.InnerActionFrameAddEasingData(builder, easingData)
    }
    return csbparser.InnerActionFrameEnd(builder)
}
func CreateEasingData(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var easeType int32
    hasEaseType := false
    if __value, ok := jValue["easeType"]; ok {
        if __value_, ok := __value.(float64); ok {
            easeType = int32(__value_)
            hasEaseType = true
        }
    }
    var points flatbuffers.UOffsetT
    hasPoints := false
    if __value, ok := jValue["points"]; ok {
        if __value_, ok := __value.([]interface{}); ok {
            vector := []flatbuffers.UOffsetT{}
            for _, v := range __value_ {
                if v_, ok := v.(map[string]interface{}); ok {
                    vector = append(vector, CreatePosition(builder, v_))
                }
            }
            csbparser.EasingDataStartPointsVector(builder, len(vector))
            for i := len(vector)-1; i >= 0; i-- {
                builder.PrependUOffsetT(vector[i])
            }
            points = builder.EndVector(len(vector))
            hasPoints = true
        }
    }
    csbparser.EasingDataStart(builder)
    if hasEaseType {
        csbparser.EasingDataAddEaseType(builder, easeType)
    }
    if hasPoints {
        csbparser.EasingDataAddPoints(builder, points)
    }
    return csbparser.EasingDataEnd(builder)
}
func CreateRotationSkew(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var rotationSkewX float32
    hasRotationSkewX := false
    if __value, ok := jValue["rotationSkewX"]; ok {
        if __value_, ok := __value.(float64); ok {
            rotationSkewX = float32(__value_)
            hasRotationSkewX = true
        }
    }
    var rotationSkewY float32
    hasRotationSkewY := false
    if __value, ok := jValue["rotationSkewY"]; ok {
        if __value_, ok := __value.(float64); ok {
            rotationSkewY = float32(__value_)
            hasRotationSkewY = true
        }
    }
    if !hasRotationSkewX && !hasRotationSkewY {
        return 0
    }
    return csbparser.CreateRotationSkew(builder, rotationSkewX, rotationSkewY)
}
func CreatePosition(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var x float32
    hasX := false
    if __value, ok := jValue["x"]; ok {
        if __value_, ok := __value.(float64); ok {
            x = float32(__value_)
            hasX = true
        }
    }
    var y float32
    hasY := false
    if __value, ok := jValue["y"]; ok {
        if __value_, ok := __value.(float64); ok {
            y = float32(__value_)
            hasY = true
        }
    }
    if !hasX && !hasY {
        return 0
    }
    return csbparser.CreatePosition(builder, x, y)
}
func CreateScale(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var scaleX float32
    hasScaleX := false
    if __value, ok := jValue["scaleX"]; ok {
        if __value_, ok := __value.(float64); ok {
            scaleX = float32(__value_)
            hasScaleX = true
        }
    }
    var scaleY float32
    hasScaleY := false
    if __value, ok := jValue["scaleY"]; ok {
        if __value_, ok := __value.(float64); ok {
            scaleY = float32(__value_)
            hasScaleY = true
        }
    }
    if !hasScaleX && !hasScaleY {
        return 0
    }
    return csbparser.CreateScale(builder, scaleX, scaleY)
}
func CreateAnchorPoint(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var scaleX float32
    hasScaleX := false
    if __value, ok := jValue["scaleX"]; ok {
        if __value_, ok := __value.(float64); ok {
            scaleX = float32(__value_)
            hasScaleX = true
        }
    }
    var scaleY float32
    hasScaleY := false
    if __value, ok := jValue["scaleY"]; ok {
        if __value_, ok := __value.(float64); ok {
            scaleY = float32(__value_)
            hasScaleY = true
        }
    }
    if !hasScaleX && !hasScaleY {
        return 0
    }
    return csbparser.CreateAnchorPoint(builder, scaleX, scaleY)
}
func CreateColor(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var a byte
    hasA := false
    if __value, ok := jValue["a"]; ok {
        if __value_, ok := __value.(float64); ok {
            a = byte(__value_)
            hasA = true
        }
    }
    var r byte
    hasR := false
    if __value, ok := jValue["r"]; ok {
        if __value_, ok := __value.(float64); ok {
            r = byte(__value_)
            hasR = true
        }
    }
    var g byte
    hasG := false
    if __value, ok := jValue["g"]; ok {
        if __value_, ok := __value.(float64); ok {
            g = byte(__value_)
            hasG = true
        }
    }
    var b byte
    hasB := false
    if __value, ok := jValue["b"]; ok {
        if __value_, ok := __value.(float64); ok {
            b = byte(__value_)
            hasB = true
        }
    }
    if !hasA && !hasR && !hasG && !hasB {
        return 0
    }
    return csbparser.CreateColor(builder, a, r, g, b)
}
func CreateColorVector(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var vectorX float32
    hasVectorX := false
    if __value, ok := jValue["vectorX"]; ok {
        if __value_, ok := __value.(float64); ok {
            vectorX = float32(__value_)
            hasVectorX = true
        }
    }
    var vectorY float32
    hasVectorY := false
    if __value, ok := jValue["vectorY"]; ok {
        if __value_, ok := __value.(float64); ok {
            vectorY = float32(__value_)
            hasVectorY = true
        }
    }
    if !hasVectorX && !hasVectorY {
        return 0
    }
    return csbparser.CreateColorVector(builder, vectorX, vectorY)
}
func CreateFlatSize(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var width float32
    hasWidth := false
    if __value, ok := jValue["width"]; ok {
        if __value_, ok := __value.(float64); ok {
            width = float32(__value_)
            hasWidth = true
        }
    }
    var height float32
    hasHeight := false
    if __value, ok := jValue["height"]; ok {
        if __value_, ok := __value.(float64); ok {
            height = float32(__value_)
            hasHeight = true
        }
    }
    if !hasWidth && !hasHeight {
        return 0
    }
    return csbparser.CreateFlatSize(builder, width, height)
}
func CreateCapInsets(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var x float32
    hasX := false
    if __value, ok := jValue["x"]; ok {
        if __value_, ok := __value.(float64); ok {
            x = float32(__value_)
            hasX = true
        }
    }
    var y float32
    hasY := false
    if __value, ok := jValue["y"]; ok {
        if __value_, ok := __value.(float64); ok {
            y = float32(__value_)
            hasY = true
        }
    }
    var width float32
    hasWidth := false
    if __value, ok := jValue["width"]; ok {
        if __value_, ok := __value.(float64); ok {
            width = float32(__value_)
            hasWidth = true
        }
    }
    var height float32
    hasHeight := false
    if __value, ok := jValue["height"]; ok {
        if __value_, ok := __value.(float64); ok {
            height = float32(__value_)
            hasHeight = true
        }
    }
    if !hasX && !hasY && !hasWidth && !hasHeight {
        return 0
    }
    return csbparser.CreateCapInsets(builder, x, y, width, height)
}
func CreateBlendFunc(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var src int32
    hasSrc := false
    if __value, ok := jValue["src"]; ok {
        if __value_, ok := __value.(float64); ok {
            src = int32(__value_)
            hasSrc = true
        }
    }
    var dst int32
    hasDst := false
    if __value, ok := jValue["dst"]; ok {
        if __value_, ok := __value.(float64); ok {
            dst = int32(__value_)
            hasDst = true
        }
    }
    if !hasSrc && !hasDst {
        return 0
    }
    return csbparser.CreateBlendFunc(builder, src, dst)
}
func CreateResourceData(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var path flatbuffers.UOffsetT
    hasPath := false
    if __value, ok := jValue["path"]; ok {
        if __value_, ok := __value.(string); ok {
            path = builder.CreateString(__value_)
            hasPath = true
        }
    }
    var plistFile flatbuffers.UOffsetT
    hasPlistFile := false
    if __value, ok := jValue["plistFile"]; ok {
        if __value_, ok := __value.(string); ok {
            plistFile = builder.CreateString(__value_)
            hasPlistFile = true
        }
    }
    var resourceType int32
    hasResourceType := false
    if __value, ok := jValue["resourceType"]; ok {
        if __value_, ok := __value.(float64); ok {
            resourceType = int32(__value_)
            hasResourceType = true
        }
    }
    csbparser.ResourceDataStart(builder)
    if hasPath {
        csbparser.ResourceDataAddPath(builder, path)
    }
    if hasPlistFile {
        csbparser.ResourceDataAddPlistFile(builder, plistFile)
    }
    if hasResourceType {
        csbparser.ResourceDataAddResourceType(builder, resourceType)
    }
    return csbparser.ResourceDataEnd(builder)
}
func CreateBlendFrame(builder *flatbuffers.Builder, jValue map[string]interface{}) flatbuffers.UOffsetT {
    var frameIndex int32
    hasFrameIndex := false
    if __value, ok := jValue["frameIndex"]; ok {
        if __value_, ok := __value.(float64); ok {
            frameIndex = int32(__value_)
            hasFrameIndex = true
        }
    }
    var tween bool
    hasTween := false
    if __value, ok := jValue["tween"]; ok {
        if __value_, ok := __value.(bool); ok {
            tween = bool(__value_)
            hasTween = true
        }
    }
    var blendFunc map[string]interface{}
    hasBlendFunc := false
    if __value, ok := jValue["blendFunc"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            blendFunc = __value_
            hasBlendFunc = true
        }
    }
    var easingData flatbuffers.UOffsetT
    hasEasingData := false
    if __value, ok := jValue["easingData"]; ok {
        if __value_, ok := __value.(map[string]interface{}); ok {
            easingData = CreateEasingData(builder, __value_)
            hasEasingData = true
        }
    }
    csbparser.BlendFrameStart(builder)
    if hasFrameIndex {
        csbparser.BlendFrameAddFrameIndex(builder, frameIndex)
    }
    if hasTween {
        csbparser.BlendFrameAddTween(builder, tween)
    }
    if hasBlendFunc {
        offset := CreateBlendFunc(builder, blendFunc)
        csbparser.BlendFrameAddBlendFunc(builder, offset)
    }
    if hasEasingData {
        csbparser.BlendFrameAddEasingData(builder, easingData)
    }
    return csbparser.BlendFrameEnd(builder)
}

