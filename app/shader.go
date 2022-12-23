package win

/*
I got this code
from a demo in g3n repo
*/

func (w *Window) AddShaders() {

	// Create custom shader
	w.app.Renderer().AddShader("shaderBricksVertex", shaderBricksVertex)
	w.app.Renderer().AddShader("shaderBricksFrag", shaderBricksFrag)
	w.app.Renderer().AddProgram("shaderBricks", "shaderBricksVertex", "shaderBricksFrag")

}

// Vertex Shader
const shaderBricksVertex = `
#include <attributes>
#include <material>
// Model uniforms
uniform mat4 ModelViewMatrix;
uniform mat3 NormalMatrix;
uniform mat4 MVP;
// Output variables for Fragment shader
out vec4 Position;
out vec3 Normal;
out vec3 CamDir;
out vec2 FragTexcoord;
out vec2 VPosition;
void main() {
    // Transform this vertex position to camera coordinates.
    Position = ModelViewMatrix * vec4(VertexPosition, 1.0);
    VPosition = VertexPosition.xy;
    // Transform this vertex normal to camera coordinates.
    Normal = normalize(NormalMatrix * VertexNormal);
    // Calculate the direction vector from the vertex to the camera
    // The camera is at 0,0,0
    CamDir = normalize(-Position.xyz);
    // Flips texture coordinate Y if requested.
    vec2 texcoord = VertexTexcoord;
	#if MAT_TEXTURES>0
    if (MatTexFlipY[0] > 0) {
        texcoord.y = 1 - texcoord.y;
    }
	#endif
    FragTexcoord = texcoord;
    gl_Position = MVP * vec4(VertexPosition, 1.0);
}
`

// Fragment Shader
const shaderBricksFrag = `
precision highp float;
// Inputs from vertex shader
in vec4 Position;       // Vertex position in camera coordinates.
in vec3 Normal;         // Vertex normal in camera coordinates.
in vec3 CamDir;         // Direction from vertex to camera
in vec2 FragTexcoord;
in vec2 VPosition;      // Vertex position in model coordinates (xy)
#include <lights>
#include <material>
#include <phong_model>
// Uniforms for configure brick pattern
uniform vec3 BrickColor;
uniform vec3 MortarColor;
uniform vec2 BrickSize;
uniform vec2 BrickPercent;
// Final fragment color
out vec4 FragColor;
void main() {
    vec2 position = VPosition / BrickSize;
    if (fract(position.y * 0.5) > 0.5) {
        position.x += 0.5;
    }
    position = fract(position);
    vec2 useBrick = step(position, BrickPercent);
    vec3 color = mix(MortarColor, BrickColor, useBrick.x * useBrick.y);
    // Combine material with brick pattern colors
    vec4 matDiffuse = vec4(color, 1.0);
    vec4 matAmbient = vec4(MatAmbientColor, MatOpacity) * vec4(color, 1.0);
    // Inverts the fragment normal if not FrontFacing
    vec3 fragNormal = Normal;
    if (!gl_FrontFacing) {
        fragNormal = -fragNormal;
    }
    // Calculates the Ambient+Diffuse and Specular colors for this fragment using the Phong model.
    vec3 Ambdiff, Spec;
    phongModel(Position, fragNormal, CamDir, vec3(matAmbient), vec3(matDiffuse), Ambdiff, Spec);
    // Final fragment color
    FragColor = min(vec4(Ambdiff + Spec, matDiffuse.a), vec4(1.0));
}
`
