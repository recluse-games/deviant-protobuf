// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: Direction.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Deviant {

  /// <summary>Holder for reflection information generated from Direction.proto</summary>
  public static partial class DirectionReflection {

    #region Descriptor
    /// <summary>File descriptor for Direction.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static DirectionReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cg9EaXJlY3Rpb24ucHJvdG8SB0RldmlhbnQqMgoJRGlyZWN0aW9uEgYKAlVQ",
            "EAASCAoERE9XThABEggKBExFRlQQAhIJCgVSSUdIVBADQglaB2RldmlhbnRi",
            "BnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Deviant.Direction), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum Direction {
    [pbr::OriginalName("UP")] Up = 0,
    [pbr::OriginalName("DOWN")] Down = 1,
    [pbr::OriginalName("LEFT")] Left = 2,
    [pbr::OriginalName("RIGHT")] Right = 3,
  }

  #endregion

}

#endregion Designer generated code
